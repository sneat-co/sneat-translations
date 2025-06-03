package main

import (
	"context"
	"fmt"
	"github.com/sneat-co/sneat-translations/trans"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"golang.org/x/text/language"
	"os"
	"sort"
	"strings"
)

func localeKeyToCode(s string) string {
	s = strings.Trim(s, `"`)
	switch len(s) {
	case 0:
		panic("empty locale code")
	case 5:
		return s
	case 4:
		return s[:2] + "-" + s[2:]
	}
	panic("unsupported locale:" + s)
}

func localeCodeToKey(s string) string {
	return strings.Replace(s, "-", "", 1)
}

func translateLocale(sourceLang language.Tag, locale string, texts []string) (translatedTexts []string, err error) {
	translatedTexts = translateMock(texts)
	//translatedTexts, err = translateByGoogle(sourceLang, locale, texts)
	return
}

func translateByGoogle(sourceLang language.Tag, locale string, texts []string) (translatedText []string, err error) {
	var targetLanguage language.Tag
	if targetLanguage, err = language.Parse(locale); err != nil {
		err = fmt.Errorf("invalid target language: %w", err)
		return
	}
	translatedText, err = translateMulti(context.Background(), sourceLang, targetLanguage, texts)
	return
}

func translateMock(texts []string) []string {
	for i, text := range texts {
		texts[i] = "TRANSLATED: " + text
	}
	return texts
}

func getKey(kv *ast.KeyValueExpr) (key string) {
	switch kvKey := kv.Key.(type) {
	case *ast.BasicLit:
		key = kvKey.Value
	case *ast.Ident:
		key = kvKey.Name
	default:
		panic("unexpected key type")
	}
	return
}

func getValue(kv *ast.KeyValueExpr) (value string, exp ast.Expr) {
	switch v := kv.Value.(type) {
	case *ast.BasicLit:
		return v.Value[1 : len(v.Value)-1], kv.Value
	case *ast.BinaryExpr:
		return getBinaryExprValue(v), kv.Value
	default:
		panic(fmt.Errorf("unexpected value type %T", kv.Value))
	}

}

func getKeyValue(kv *ast.KeyValueExpr) (key, value string) {
	key = getKey(kv)
	value, _ = getValue(kv)
	return
}

func getTextToTranslate(baseLocale string, translations *ast.CompositeLit) (textToTranslate string, isQuoted bool) {
	for _, elt := range translations.Elts {
		if innerKV, ok := elt.(*ast.KeyValueExpr); ok {
			key := getKey(innerKV)
			locale := localeKeyToCode(key)
			if locale == baseLocale {
				textToTranslate, _ = getValue(innerKV)
				isQuoted = strings.HasPrefix(innerKV.Value.(*ast.BasicLit).Value, `"`)
				return
			}
		}
	}
	return
}

func getIsAlreadyPresent(translations *ast.CompositeLit, locale string) bool {
	for _, innerElt := range translations.Elts {
		if innerKV, ok := innerElt.(*ast.KeyValueExpr); ok {
			key := getKey(innerKV)
			if localeKeyToCode(key) == locale {
				return true
			}
		}
	}
	return false
}

func getCurrentValue(translations *ast.CompositeLit, locale string) (value string, expr ast.Expr) {
	for _, innerElt := range translations.Elts {
		if innerKV, ok := innerElt.(*ast.KeyValueExpr); ok {
			key := getKey(innerKV)
			if localeKeyToCode(key) == locale {
				return getValue(innerKV)
			}
		}
	}
	return "", nil
}

func getBinaryExprValue(e *ast.BinaryExpr) string {
	var xs, ys string
	switch x := e.X.(type) {
	case *ast.BasicLit:
		xs = x.Value[1 : len(x.Value)-1]
	case *ast.BinaryExpr:
		xs = getBinaryExprValue(x)
	}
	switch y := e.Y.(type) {
	case *ast.BasicLit:
		ys = y.Value[1 : len(y.Value)-1]
	case *ast.BinaryExpr:
		ys = getBinaryExprValue(y)
	}
	if e.Op != token.ADD {
		panic("unsupported binary expressions")
	}
	return xs + ys
}

func translateMissing(baseLocale string) error {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, "../trans/translations.go", nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	// Create CommentMap
	//commentMap := ast.NewCommentMap(fileSet, node, node.Comments)

	toTranslate := make(map[string]map[string]string)
	isQuoted := make(map[string]bool)

	// Walk through AST
	ast.Inspect(node, func(n ast.Node) bool {
		kv, ok := n.(*ast.CompositeLit)
		if !ok {
			return true
		}

		// translate map[locale]map[id]{input|output}
		for _, elt := range kv.Elts {
			outerKV, ok := elt.(*ast.KeyValueExpr)
			if !ok {
				continue
			}

			translationID := getKey(outerKV)

			innerMap, ok := outerKV.Value.(*ast.CompositeLit)
			if !ok {
				continue
			}

			var textToTranslate string
			textToTranslate, isQuoted[translationID] = getTextToTranslate(baseLocale, innerMap)

			if textToTranslate != "" {
				for _, supportedLocale := range trans.SupportedLocales {
					currentValue, valueExpr := getCurrentValue(innerMap, supportedLocale)
					if currentValue == "" && !strings.HasSuffix(translationID, "S1NGL") {
						if hasTranslatableTextParts(textToTranslate) {
							tk, ok := toTranslate[supportedLocale]
							if !ok {
								tk = make(map[string]string)
								toTranslate[supportedLocale] = tk
							}
							tk[translationID] = textToTranslate
						} else {
							if valueExpr == nil {
								innerMap.Elts = append(innerMap.Elts, &ast.KeyValueExpr{
									Key:   &ast.Ident{Name: localeCodeToKey(supportedLocale)},
									Value: valueExpr,
								})
							} else {
								panic(fmt.Errorf("value of %s.%s is empty string", translationID, supportedLocale))
							}
						}
					}
				}
			}
		}

		return true
	})

	type kv struct {
		key   string
		value string
	}

	translated := make(map[string]map[string]string)

	sourceLang, err := language.Parse("en")
	if err != nil {
		return err
	}

	for locale, texts := range toTranslate {
		textsToTranslate := make([]string, 0, len(texts))
		kvs := make([]kv, 0, len(texts))
		for id, text := range texts {
			kvs = append(kvs, kv{key: id, value: text})
			textsToTranslate = append(textsToTranslate, text)
		}
		var translatedTexts []string
		if translatedTexts, err = translateLocale(sourceLang, locale, textsToTranslate); err != nil {
			return err
		}
		for i, v := range translatedTexts {
			translations := translated[kvs[i].key]
			if translations == nil {
				translations = make(map[string]string)
				translated[kvs[i].key] = translations
			}
			translations[locale] = v
		}
	}

	translated = make(map[string]map[string]string)

	ast.Inspect(node, func(n ast.Node) bool {
		kv, ok := n.(*ast.CompositeLit)
		if !ok {
			return true
		}

		// translate map[locale]map[id]{input|output}
		for _, elt := range kv.Elts {
			outerKV, ok := elt.(*ast.KeyValueExpr)
			if !ok {
				continue
			}

			innerMap, ok := outerKV.Value.(*ast.CompositeLit)
			if !ok {
				continue
			}

			translationID := getKey(outerKV)

			translations := translated[translationID]

			for locale, translation := range translations {
				var value string
				if isQuoted[translationID] {
					value = `"` + translation + `"`
				} else {
					value = "`" + translation + "`"
				}
				newKV := &ast.KeyValueExpr{
					Key: &ast.BasicLit{
						Kind:  token.STRING,
						Value: localeCodeToKey(locale),
					},
					Value: &ast.BasicLit{
						Kind:  token.STRING,
						Value: value,
					},
				}

				//commentGroup := &ast.CommentGroup{
				//	List: []*ast.Comment{
				//		{Text: "// comment next to key " + localeCodeToKey(locale)},
				//	},
				//}
				//node.Comments = append(node.Comments, commentGroup)
				//commentMap[newKV] = []*ast.CommentGroup{commentGroup}
				innerMap.Elts = append(innerMap.Elts, newKV)
			}

			sort.Slice(innerMap.Elts, func(i, j int) bool {
				kI := getKey(innerMap.Elts[i].(*ast.KeyValueExpr))
				kJ := getKey(innerMap.Elts[j].(*ast.KeyValueExpr))
				return kI < kJ
			})

			rebuildKeyValues := func() (newElements []ast.Expr) {
				//baseText, _ := getTextToTranslate(baseLocale, innerMap)
				newElements = make([]ast.Expr, 0, len(innerMap.Elts))
				for _, el := range innerMap.Elts {
					kve := el.(*ast.KeyValueExpr)
					key := getKey(kve)
					// Adding new line as separate expression does not work as adds comma before KV pair
					//newElements = append(newElements, &ast.BasicLit{Value: "\n"})

					// This is a dirty hack trying to make map keys to be on a new line
					keyValue := "\n\t\t" + key
					//localeCode := localeKeyToCode(key)
					//elValue := getValue(kve)
					//if elValue == "" || !hasTranslatableTextParts(elValue) && elValue == baseText && localeCode != "en-UK" {
					//	continue
					//}
					kve = &ast.KeyValueExpr{
						Key:   &ast.BasicLit{Kind: token.STRING, Value: keyValue},
						Value: kve.Value,
					}
					newElements = append(newElements, kve)

				}
				return
			}
			innerMap.Elts = rebuildKeyValues()

			// Force separate lines by assigning new positions
			//for i, el := range newElements {
			//	kve := el.(*ast.KeyValueExpr)
			//	// Give each element a new position in the FileSet
			//	pos := token.Pos(i*10 + 1) // Arbitrary increasing offset
			//	fileSet.AddFile("dummy.go", fileSet.Base(), 1000).AddLine(int(pos))
			//	kve.Key.Pos() // this line ensures keys are kept as-is
			//}

			//innerMap.Elts = nil
			//for _, newEl := range newElements {
			//	innerMap.Elts = append(innerMap.Elts, newEl)
			//}
		}

		return true
	})

	// Save result to file
	outFile, err := os.Create("../trans/translations.go")
	if err != nil {
		panic(err)
	}
	defer func(outFile *os.File) {
		_ = outFile.Close()
	}(outFile)

	//cfg := &printer.Config{
	//	Mode: printer.UseSpaces | printer.TabIndent,
	//	//Tabwidth: 8,
	//}
	//if err = cfg.Fprint(outFile, fileSet, node); err != nil {
	//	panic(err)
	//}
	if err = format.Node(outFile, fileSet, node); err != nil {
		panic(err)
	}
	return nil
}

func hasTranslatableTextParts(s string) bool {
	s = string(trans.ReVars.ReplaceAll([]byte(s), []byte{}))
	s = string(trans.ReTags.ReplaceAll([]byte(s), []byte{}))
	s = strings.ReplaceAll(s, " ", "")
	return s != ""
}

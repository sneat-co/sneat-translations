package main

import (
	"context"
	"fmt"
	"github.com/sneat-co/sneat-translations/trans"
	"go/ast"
	"go/parser"
	"go/printer"
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

// toExprList ensures each entry is an Expr with newline potential
func toExprList(kvs ...*ast.KeyValueExpr) []ast.Expr {
	out := make([]ast.Expr, len(kvs))
	for i, kv := range kvs {
		out[i] = kv
	}
	return out
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

func getKeyValue(kv *ast.KeyValueExpr) (key, value string) {
	switch kvKey := kv.Key.(type) {
	case *ast.BasicLit:
		key = kvKey.Value
	case *ast.Ident:
		key = kvKey.Name
	default:
		panic("unexpected key type")
	}
	return key, kv.Value.(*ast.BasicLit).Value
}

func translateMissing(baseLocale string) error {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, "../trans/translations.go", nil, parser.AllErrors|parser.ParseComments)
	if err != nil {
		panic(err)
	}

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

			outerKey, ok := outerKV.Key.(*ast.BasicLit)
			if !ok {
				continue
			}

			translationID := outerKey.Value

			innerMap, ok := outerKV.Value.(*ast.CompositeLit)
			if !ok {
				continue
			}

			var textToTranslate string

			for _, innerElt := range innerMap.Elts {
				if innerKV, ok := innerElt.(*ast.KeyValueExpr); ok {
					key, value := getKeyValue(innerKV)
					locale := localeKeyToCode(key)
					if locale == baseLocale {
						s := strings.TrimSpace(value)
						isQuoted[translationID] = strings.HasPrefix(s, `"`)
						textToTranslate = s[1 : len(s)-1]
						break
					}
				}
			}

			for _, supportedLocale := range trans.SupportedLocales {
				alreadyPresent := false

				for _, innerElt := range innerMap.Elts {
					if innerKV, ok := innerElt.(*ast.KeyValueExpr); ok {
						key, _ := getKeyValue(innerKV)
						locale := localeKeyToCode(key)
						if locale == supportedLocale {
							alreadyPresent = true
							break
						}
					}
				}

				if !alreadyPresent && textToTranslate != "" {
					tk, ok := toTranslate[supportedLocale]
					if !ok {
						tk = make(map[string]string)
						toTranslate[supportedLocale] = tk
					}
					outerKey, ok := outerKV.Key.(*ast.BasicLit)
					if !ok || outerKey.Kind != token.STRING {
						continue
					}
					tk[outerKey.Value] = textToTranslate
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
			//if len(toTranslate) > 10 {
			//	break
			//}
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

			outerKey, ok := outerKV.Key.(*ast.BasicLit)
			if !ok || outerKey.Kind != token.STRING {
				continue
			}

			translations, ok := translated[outerKey.Value]
			if !ok {
				continue
			}
			for locale, translation := range translations {
				var value string
				if isQuoted[outerKey.Value] {
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
				innerMap.Elts = append(innerMap.Elts, toExprList(newKV)...)
			}

			sort.Slice(innerMap.Elts, func(i, j int) bool {
				kI, _ := getKeyValue(innerMap.Elts[i].(*ast.KeyValueExpr))
				kJ, _ := getKeyValue(innerMap.Elts[j].(*ast.KeyValueExpr))
				return kI < kJ
			})
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

	cfg := &printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	if err = cfg.Fprint(outFile, fileSet, node); err != nil {
		panic(err)
	}
	//// Format AST and write
	//if err = format.Node(outFile, fileSet, node); err != nil {
	//	panic(err)
	//}
	return nil
}

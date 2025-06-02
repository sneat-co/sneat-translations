package main

import (
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"golang.org/x/text/language"
	"os"
	"strings"

	"github.com/sneat-co/sneat-translations/trans"
)

const baseLocale = "en-UK"

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}

func run() error {
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

			innerMap, ok := outerKV.Value.(*ast.CompositeLit)
			if !ok {
				continue
			}

			var textToTranslate string

			for _, innerElt := range innerMap.Elts {
				if innerKV, ok := innerElt.(*ast.KeyValueExpr); ok {
					if innerKey, ok := innerKV.Key.(*ast.BasicLit); ok {
						if locale := strings.Trim(innerKey.Value, `"`); locale == baseLocale {
							s := strings.TrimSpace(innerKV.Value.(*ast.BasicLit).Value)
							isQuoted[outerKey.Value] = strings.HasPrefix(s, `"`)
							textToTranslate = s[1 : len(s)-1]
						}
					}
				}
			}

			for _, supportedLocale := range trans.SupportedLocales {
				alreadyPresent := false

				for _, innerElt := range innerMap.Elts {
					if innerKV, ok := innerElt.(*ast.KeyValueExpr); ok {
						if key, ok := innerKV.Key.(*ast.BasicLit); ok {
							if locale := strings.Trim(key.Value, `"`); locale == supportedLocale {
								alreadyPresent = true
							}
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
						Value: fmt.Sprintf("\n\t\t"+`"%s"`, locale),
					},
					Value: &ast.BasicLit{
						Kind:  token.STRING,
						Value: value + ",\n",
					},
				}
				innerMap.Elts = append(innerMap.Elts, newKV)
			}
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

	//cfg := &printer.Config{Mode: printer.UseSpaces | printer.TabIndent, Tabwidth: 8}
	//if err = cfg.Fprint(outFile, fileSet, node); err != nil {
	//	panic(err)
	//}
	// Format AST and write
	if err = format.Node(outFile, fileSet, node); err != nil {
		panic(err)
	}
	return nil
}

func translateLocale(sourceLang language.Tag, locale string, texts []string) (translatedText []string, err error) {
	var targetLang language.Tag
	if targetLang, err = language.Parse(locale); err != nil {
		err = fmt.Errorf("invalid target language: %w", err)
		return
	}
	translatedText, err = translateMulti(context.Background(), sourceLang, targetLang, texts)
	return
}

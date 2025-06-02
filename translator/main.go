package main

import (
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"strings"
)

func main() {
	fileSet := token.NewFileSet()
	node, err := parser.ParseFile(fileSet, "../trans/translations.go", nil, parser.AllErrors)
	if err != nil {
		panic(err)
	}

	// Walk through AST
	ast.Inspect(node, func(n ast.Node) bool {
		kv, ok := n.(*ast.CompositeLit)
		if !ok {
			return true
		}

		for _, elt := range kv.Elts {
			outerKV, ok := elt.(*ast.KeyValueExpr)
			if !ok {
				continue
			}

			innerMap, ok := outerKV.Value.(*ast.CompositeLit)
			if !ok {
				continue
			}

			for _, innerElt := range innerMap.Elts {
				innerKV, ok := innerElt.(*ast.KeyValueExpr)
				if !ok {
					continue
				}

				if val, ok := innerKV.Value.(*ast.BasicLit); ok && val.Kind == token.STRING {
					val.Value = `"` + strings.ToUpper(strings.Trim(val.Value, `"`)) + `"`
				}
			}
		}

		return true
	})

	// Save result to file
	outFile, err := os.Create("translations_modified.go")
	if err != nil {
		panic(err)
	}
	defer func(outFile *os.File) {
		_ = outFile.Close()
	}(outFile)

	_ = printer.Fprint(outFile, fileSet, node)
}

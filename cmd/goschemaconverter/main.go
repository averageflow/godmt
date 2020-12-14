package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../tests/data/examplestructs/Example.go", nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	for i := range f.Scope.Objects {
		token := f.Scope.Objects[i].Decl
		tokenType := fmt.Sprintf("%T", token)

		if tokenType == "*ast.ValueSpec" {
			item := token.(*ast.ValueSpec).Values[0]

			fmt.Printf("item: %+v\n", item)
			itemType := fmt.Sprintf("%T", item)

			if itemType == "*ast.BasicLit" {
				fmt.Printf("%+v\n", item.(*ast.BasicLit).Kind)
			} else if itemType == "*ast.Ident" {
				fmt.Printf("%+v\n", item.(*ast.Ident).String())
			}
		} else if tokenType == "*ast.TypeSpec" {
			item := token.(*ast.TypeSpec)
			fmt.Printf("TOKEN IS: %v", item)
		}

	}

	fmt.Print("EXITING")
}

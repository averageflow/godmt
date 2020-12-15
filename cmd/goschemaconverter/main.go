package main

import (
	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../tests/data/exampleconstants/Example.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	syntaxtree.WalkSyntaxTree(f)
}

/*
	var identifierItems []ast.BasicLit
	var typeItems []ast.TypeSpec
	var basicTypeItems []ast.Ident

	for i := range f.Scope.Objects {
		scannedToken := f.Scope.Objects[i].Decl

		switch scannedToken.(type) {
		case *ast.ValueSpec:
			item := scannedToken.(*ast.ValueSpec).Values[0]

			switch item.(type) {
			case *ast.BasicLit:
				identifierItems = append(identifierItems, *item.(*ast.BasicLit))
				//fmt.Printf("%+v\n", item.(*ast.BasicLit).Kind)
				break
			case *ast.Ident:
				basicTypeItems = append(basicTypeItems, *item.(*ast.Ident))
				//fmt.Printf("%+v\n", item.(*ast.Ident).String())
				break
			}

			break
		case *ast.TypeSpec:
			item := scannedToken.(*ast.TypeSpec)
			//fmt.Printf("TOKEN IS: %v", item)
			typeItems = append(typeItems, *item)
			break
		}

	}

	for i := range typeItems {
		structName := typeItems[i].Name.Name
		structFields := reflect.ValueOf(typeItems[i].Type).Elem().FieldByName("Fields").Elem()

		structFieldList := structFields.FieldByName("List").Interface().([]*ast.Field)

		for j := range structFieldList {
			fmt.Printf("Name: %+v\n", structName)

			subStructFields := reflect.ValueOf(structFieldList[j].Type).Elem().FieldByName("Fields").Interface().(*ast.FieldList).List
			for k := range subStructFields {
				fmt.Printf("name: %v | kind: %v | tag: %v\n", subStructFields[k].Names[0].Name, subStructFields[k].Tag.Kind, subStructFields[k].Tag.Value)
			}
		}

	}*/

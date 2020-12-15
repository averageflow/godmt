package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"reflect"
)

type ParsedEnum struct {
	Name  string
	Kind  string
	Value interface{}
}

type visitor int

func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	switch d := n.(type) {

	case *ast.Ident:
		if d.Obj == nil {
			return v + 1
		}

		if d.Obj.Kind == ast.Typ {

		}

		if d.Obj.Kind == ast.Con || d.Obj.Kind == ast.Var {
			v.ParseConstantsAndVariables(d)
		}

		break
	}
	//fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}

func (v visitor) ParseConstantsAndVariables(d *ast.Ident) {
	objectValues := reflect.ValueOf(d.Obj.Decl).Elem().FieldByName("Values")

	values := objectValues.Interface().([]ast.Expr)

	for i := range values {
		switch values[i].(type) {
		case *ast.BasicLit:
			item := values[i].(*ast.BasicLit)
			itemType := fmt.Sprintf("%T", item.Value)
			if item.Kind == token.INT {
				itemType = "int"
			}

			fmt.Printf("%+v\n", ParsedEnum{
				Name:  d.Name,
				Kind:  itemType,
				Value: item.Value,
			})
		case *ast.Ident:
			item := values[i].(*ast.Ident)
			if item.Name == "true" || item.Name == "false" {
				fmt.Printf("%+v\n", ParsedEnum{
					Name:  d.Name,
					Kind:  "bool",
					Value: item.Name,
				})
			}
		case *ast.CompositeLit:
			item := values[i].(*ast.CompositeLit)
			switch item.Type.(type) {
			case *ast.MapType:

				mapElements := reflect.ValueOf(item.Elts).Interface().([]ast.Expr)
				cleanMap := make(map[interface{}]interface{})

				for j := range mapElements {
					rawKey := reflect.ValueOf(mapElements[j]).Elem().FieldByName("Key").Interface().(*ast.BasicLit)
					rawValue := reflect.ValueOf(mapElements[j]).Elem().FieldByName("Value").Interface().(*ast.BasicLit)
					cleanMap[rawKey.Value] = rawValue.Value
				}

				fmt.Printf("%+v\n", ParsedEnum{
					Name:  d.Name,
					Kind:  fmt.Sprintf("map[%s]%s", item.Type.(*ast.MapType).Key.(*ast.Ident).Name, item.Type.(*ast.MapType).Value.(*ast.Ident).Name),
					Value: cleanMap,
				})
			}
		}
	}
}

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "../../tests/data/examplevars/Example.go", nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}
	var v visitor

	ast.Walk(v, f)

	fmt.Print("EXITING")
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

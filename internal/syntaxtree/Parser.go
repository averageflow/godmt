package syntaxtree

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

func parseConstantsAndVariables(d *ast.Ident) {
	objectValues := reflect.ValueOf(d.Obj.Decl).Elem().FieldByName("Values")

	values := objectValues.Interface().([]ast.Expr)

	for i := range values {
		switch values[i].(type) {
		case *ast.BasicLit:
			item := values[i].(*ast.BasicLit)
			itemType := fmt.Sprintf("%T", item.Value)
			if item.Kind == token.INT {
				itemType = "int64"
			} else if item.Kind == token.FLOAT {
				itemType = "float64"
			}

			var doc []string

			rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)
			if rawDecl.Doc != nil {
				for i := range rawDecl.Doc.List {
					doc = append(doc, rawDecl.Doc.List[i].Text)
				}
			}

			fmt.Printf("%+v\n", RawScannedType{
				Name:         d.Name,
				Kind:         itemType,
				Value:        item.Value,
				Doc:          doc,
				InternalType: ConstType,
			})
		case *ast.Ident:
			item := values[i].(*ast.Ident)

			var doc []string

			rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)
			if rawDecl.Doc != nil {
				for i := range rawDecl.Doc.List {
					doc = append(doc, rawDecl.Doc.List[i].Text)
				}

			}

			if item.Name == "true" || item.Name == "false" {
				fmt.Printf("%+v\n", RawScannedType{
					Name:         d.Name,
					Kind:         "bool",
					Value:        item.Name,
					Doc:          doc,
					InternalType: ConstType,
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

				fmt.Printf("%+v\n", RawScannedType{
					Name:  d.Name,
					Kind:  fmt.Sprintf("map[%s]%s", item.Type.(*ast.MapType).Key.(*ast.Ident).Name, item.Type.(*ast.MapType).Value.(*ast.Ident).Name),
					Value: cleanMap,
				})
			}
		}
	}
}

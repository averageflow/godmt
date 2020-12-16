package syntaxtree

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

func parseStruct(d *ast.Ident) []ScannedStruct {
	var result []ScannedStruct

	structTypes := reflect.ValueOf(d.Obj.Decl).Elem().FieldByName("Type")
	if !structTypes.IsValid() {
		return result
	}

	fields := structTypes.Interface().(*ast.StructType)
	fieldList := fields.Fields.List

	var rawScannedFields []ScannedStructField

	for i := range fieldList {
		switch fieldList[i].Type.(type) {
		case *ast.Ident:
			// Basic type of a field inside a struct
			if fieldList[i].Names != nil {
				fieldType := reflect.ValueOf(fieldList[i].Type).Elem().FieldByName("Name")

				rawScannedFields = append(rawScannedFields, ScannedStructField{
					Doc:  extractComments(fieldList[i].Doc),
					Name: fieldList[i].Names[0].Name,
					Kind: fieldType.Interface().(string),
					Tag:  fieldList[i].Tag.Value,
				})
			} else {
				// Struct inside a struct
				fieldType := reflect.ValueOf(fieldList[i].Type).Elem().FieldByName("Obj").Interface().(*ast.Object)
				tag := fieldList[i].Tag

				var tagValue string
				if tag != nil {
					tagValue = tag.Value
				}

				rawScannedFields = append(rawScannedFields, ScannedStructField{
					Doc:  nil,
					Name: fieldType.Name,
					Kind: "struct",
					Tag:  tagValue,
				})
			}

		case *ast.StructType:
			fmt.Println("TODO: Support nested structs!")
			break
		}
	}

	result = append(result, ScannedStruct{
		Doc:          nil,
		Name:         d.Name,
		Fields:       rawScannedFields,
		InternalType: StructType,
	})

	return result
}

func parseConstantsAndVariables(d *ast.Ident) []ScannedType {
	var result []ScannedType

	objectValues := reflect.ValueOf(d.Obj.Decl).Elem().FieldByName("Values")
	if !objectValues.IsValid() {
		return result
	}

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

			rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)

			result = append(result, ScannedType{
				Name:         d.Name,
				Kind:         itemType,
				Value:        item.Value,
				Doc:          extractComments(rawDecl.Doc),
				InternalType: ConstType,
			})
		case *ast.Ident:
			item := values[i].(*ast.Ident)

			rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)

			if item.Name == "true" || item.Name == "false" {
				result = append(result, ScannedType{
					Name:         d.Name,
					Kind:         "bool",
					Value:        item.Name,
					Doc:          extractComments(rawDecl.Doc),
					InternalType: ConstType,
				})
			}
		case *ast.CompositeLit:
			item := values[i].(*ast.CompositeLit)
			switch item.Type.(type) {
			case *ast.MapType:
				mapElements := reflect.ValueOf(item.Elts).Interface().([]ast.Expr)
				cleanMap := make(map[string]string)

				for j := range mapElements {
					rawKey := reflect.ValueOf(mapElements[j]).Elem().FieldByName("Key").Interface().(*ast.BasicLit)
					rawValue := reflect.ValueOf(mapElements[j]).Elem().FieldByName("Value").Interface().(*ast.BasicLit)
					cleanMap[fmt.Sprintf("%v", rawKey.Value)] = rawValue.Value
				}

				var doc []string

				rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)
				if rawDecl.Doc != nil {
					for i := range rawDecl.Doc.List {
						doc = append(doc, rawDecl.Doc.List[i].Text)
					}
				}

				result = append(result, ScannedType{
					Name: d.Name,
					Kind: fmt.Sprintf(
						"map[%s]%s",
						getMapValueType(item.Type.(*ast.MapType).Key),
						getMapValueType(item.Type.(*ast.MapType).Value),
					),
					Value:        cleanMap,
					InternalType: MapType,
					Doc:          doc,
				})
			case *ast.ArrayType:
				sliceType := getMapValueType(item.Type.(*ast.ArrayType).Elt)

				result = append(result, ScannedType{
					Name:         d.Name,
					Kind:         fmt.Sprintf("[]%s", sliceType),
					Value:        extractSliceValues(item.Elts),
					Doc:          nil,
					InternalType: SliceType,
				})
			}
		}
	}

	return result
}

func extractComments(rawCommentGroup *ast.CommentGroup) []string {
	var result []string
	if rawCommentGroup == nil {
		return result
	}

	commentList := rawCommentGroup.List
	for i := range commentList {
		result = append(result, commentList[i].Text)
	}

	return result
}

func getMapValueType(item ast.Expr) string {
	switch item.(type) {
	case *ast.Ident:
		return item.(*ast.Ident).Name
	default:
		return "interface{}"
	}
}

func extractSliceValues(items []ast.Expr) []string {
	var result []string
	for i := range items {
		result = append(result, items[i].(*ast.BasicLit).Value)
	}
	return result
}

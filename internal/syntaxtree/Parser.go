package syntaxtree

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

	"github.com/averageflow/goschemaconverter/pkg/syntaxtreeparser"
)

func basicTypeLiteralParser(d *ast.Ident, item *ast.BasicLit) syntaxtreeparser.ScannedType {
	itemType := fmt.Sprintf("%T", item.Value)

	if item.Kind == token.INT {
		itemType = "int64"
	} else if item.Kind == token.FLOAT {
		itemType = "float64"
	}

	rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)

	return syntaxtreeparser.ScannedType{
		Name:         d.Name,
		Kind:         itemType,
		Value:        item.Value,
		Doc:          syntaxtreeparser.ExtractComments(rawDecl.Doc),
		InternalType: syntaxtreeparser.ConstType,
	}
}

func identifierParser(d *ast.Ident, item *ast.Ident) *syntaxtreeparser.ScannedType {
	rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)

	if item.Name == "true" || item.Name == "false" {
		return &syntaxtreeparser.ScannedType{
			Name:         d.Name,
			Kind:         "bool",
			Value:        item.Name,
			Doc:          syntaxtreeparser.ExtractComments(rawDecl.Doc),
			InternalType: syntaxtreeparser.ConstType,
		}
	}
	return nil
}

func compositeLiteralMapParser() {

}

func compositeLiteralSliceParser() {

}

func parseStruct(d *ast.Ident) []syntaxtreeparser.ScannedStruct {
	var result []syntaxtreeparser.ScannedStruct

	structTypes := reflect.ValueOf(d.Obj.Decl).Elem().FieldByName("Type")
	if !structTypes.IsValid() {
		return result
	}

	switch structTypes.Interface().(type) {
	case *ast.StructType:
		fields := structTypes.Interface().(*ast.StructType)
		fieldList := fields.Fields.List

		var rawScannedFields []syntaxtreeparser.ScannedStructField

		for i := range fieldList {
			switch fieldList[i].Type.(type) {
			case *ast.Ident:
				// Basic type of a field inside a struct
				if fieldList[i].Names != nil {
					fieldType := reflect.ValueOf(fieldList[i].Type).Elem().FieldByName("Name")

					tag := fieldList[i].Tag

					var tagValue string
					if tag != nil {
						tagValue = tag.Value
					}

					rawScannedFields = append(rawScannedFields, syntaxtreeparser.ScannedStructField{
						Doc:  syntaxtreeparser.ExtractComments(fieldList[i].Doc),
						Name: fieldList[i].Names[0].Name,
						Kind: fieldType.Interface().(string),
						Tag:  tagValue,
					})
				} else {
					// Struct inside a struct
					fieldType := reflect.ValueOf(fieldList[i].Type).Elem().FieldByName("Obj").Interface().(*ast.Object)
					tag := fieldList[i].Tag

					var tagValue string
					if tag != nil {
						tagValue = tag.Value
					}

					rawScannedFields = append(rawScannedFields, syntaxtreeparser.ScannedStructField{
						Doc:  nil,
						Name: fieldType.Name,
						Kind: "struct",
						Tag:  tagValue,
					})
				}
			case *ast.StructType:
				fmt.Println("TODO: Support nested structs!")
				break
			case *ast.SelectorExpr:
				// imported entity
				fieldType := reflect.ValueOf(fieldList[i].Type).Interface().(*ast.SelectorExpr)

				tag := fieldList[i].Tag

				var tagValue string
				if tag != nil {
					tagValue = tag.Value
				}

				packageName := fmt.Sprintf("%s", reflect.ValueOf(fieldType.X).Elem().FieldByName("Name"))
				rawScannedFields = append(rawScannedFields, syntaxtreeparser.ScannedStructField{
					Doc:  nil,
					Name: fieldList[i].Names[0].Name,
					Kind: fieldType.Sel.Name,
					Tag:  tagValue,
					ImportDetails: &syntaxtreeparser.ImportedEntityDetails{
						EntityName:  fieldType.Sel.Name,
						PackageName: packageName,
					},
				})
			}
		}
		result = append(result, syntaxtreeparser.ScannedStruct{
			Doc:          nil,
			Name:         d.Name,
			Fields:       rawScannedFields,
			InternalType: syntaxtreeparser.StructType,
		})
	}

	return result
}

func parseConstantsAndVariables(d *ast.Ident) []syntaxtreeparser.ScannedType {
	var result []syntaxtreeparser.ScannedType

	objectValues := reflect.ValueOf(d.Obj.Decl).Elem().FieldByName("Values")
	if !objectValues.IsValid() {
		return result
	}

	values := objectValues.Interface().([]ast.Expr)

	for i := range values {
		switch values[i].(type) {
		case *ast.BasicLit:
			item := values[i].(*ast.BasicLit)
			parsed := basicTypeLiteralParser(d, item)
			result = append(result, parsed)

		case *ast.Ident:
			item := values[i].(*ast.Ident)
			parsed := identifierParser(d, item)
			if parsed != nil {
				result = append(result, *parsed)
			}

		case *ast.CompositeLit:
			item := values[i].(*ast.CompositeLit)
			switch item.Type.(type) {
			case *ast.MapType:
				mapElements := reflect.ValueOf(item.Elts).Interface().([]ast.Expr)
				cleanMap := make(map[string]string)

				for j := range mapElements {
					rawKey := reflect.ValueOf(mapElements[j]).Elem().FieldByName("Key")
					switch rawKey.Interface().(type) {
					case *ast.BasicLit:
						rawValue := reflect.ValueOf(mapElements[j]).Elem().FieldByName("Value").Interface().(*ast.BasicLit)
						cleanMap[fmt.Sprintf("%v", rawKey.Interface().(*ast.BasicLit).Value)] = rawValue.Value
					}
				}

				var doc []string

				rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)
				if rawDecl.Doc != nil {
					for i := range rawDecl.Doc.List {
						doc = append(doc, rawDecl.Doc.List[i].Text)
					}
				}

				result = append(result, syntaxtreeparser.ScannedType{
					Name: d.Name,
					Kind: fmt.Sprintf(
						"map[%s]%s",
						syntaxtreeparser.GetMapValueType(item.Type.(*ast.MapType).Key),
						syntaxtreeparser.GetMapValueType(item.Type.(*ast.MapType).Value),
					),
					Value:        cleanMap,
					InternalType: syntaxtreeparser.MapType,
					Doc:          doc,
				})
			case *ast.ArrayType:
				sliceType := syntaxtreeparser.GetMapValueType(item.Type.(*ast.ArrayType).Elt)
				rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)

				result = append(result, syntaxtreeparser.ScannedType{
					Name:         d.Name,
					Kind:         fmt.Sprintf("[]%s", sliceType),
					Value:        syntaxtreeparser.ExtractSliceValues(item.Elts),
					Doc:          syntaxtreeparser.ExtractComments(rawDecl.Doc),
					InternalType: syntaxtreeparser.SliceType,
				})
			}
		}
	}

	return result
}

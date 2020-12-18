package syntaxtree

import (
	"fmt"
	"go/ast"
	"reflect"

	"github.com/averageflow/goschemaconverter/pkg/syntaxtreeparser"
)

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
				parsed := syntaxtreeparser.SimpleStructFieldParser(fieldList[i])
				rawScannedFields = append(rawScannedFields, parsed)
			case *ast.StructType:
				fmt.Println("TODO: Support nested structs!")
				break
			case *ast.SelectorExpr:
				parsed := syntaxtreeparser.ImportedStructFieldParser(fieldList[i])
				rawScannedFields = append(rawScannedFields, parsed)
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
			parsed := syntaxtreeparser.BasicTypeLiteralParser(d, item)
			result = append(result, parsed)

		case *ast.Ident:
			item := values[i].(*ast.Ident)
			parsed := syntaxtreeparser.IdentifierParser(d, item)
			if parsed != nil {
				result = append(result, *parsed)
			}

		case *ast.CompositeLit:
			item := values[i].(*ast.CompositeLit)
			switch item.Type.(type) {
			case *ast.MapType:
				mapElements := reflect.ValueOf(item.Elts).Interface().([]ast.Expr)
				parsed := syntaxtreeparser.CompositeLiteralMapParser(d, mapElements, item)
				result = append(result, parsed)
			case *ast.ArrayType:
				sliceType := syntaxtreeparser.GetMapValueType(item.Type.(*ast.ArrayType).Elt)
				parsed := syntaxtreeparser.CompositeLiteralSliceParser(d, sliceType, item)
				result = append(result, parsed)
			}
		}
	}

	return result
}

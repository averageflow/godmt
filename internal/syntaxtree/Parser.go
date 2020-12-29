package syntaxtree

import (
	"go/ast"
	"reflect"

	"github.com/averageflow/godmt/pkg/godmt"
)

func parseStruct(d *ast.Ident) []godmt.ScannedStruct {
	var result []godmt.ScannedStruct

	structTypes := reflect.ValueOf(d.Obj.Decl).Elem().FieldByName("Type")
	if !structTypes.IsValid() {
		return result
	}

	switch structTypes.Interface().(type) {
	case *ast.StructType:
		fields := structTypes.Interface().(*ast.StructType)
		fieldList := fields.Fields.List

		var rawScannedFields []godmt.ScannedStructField

		for i := range fieldList {
			switch fieldList[i].Type.(type) {
			case *ast.Ident:
				parsed := godmt.SimpleStructFieldParser(fieldList[i])
				rawScannedFields = append(rawScannedFields, parsed)
			case *ast.StructType:
				// fmt.Println("TODO: Support nested structs!")
				break
			case *ast.SelectorExpr:
				parsed := godmt.ImportedStructFieldParser(fieldList[i])
				rawScannedFields = append(rawScannedFields, parsed)
			}
		}

		result = append(result, godmt.ScannedStruct{
			Doc:          nil,
			Name:         d.Name,
			Fields:       rawScannedFields,
			InternalType: godmt.StructType,
		})

	default:
		break
	}

	return result
}

func parseConstantsAndVariables(d *ast.Ident) []godmt.ScannedType {
	var result []godmt.ScannedType

	objectValues := reflect.ValueOf(d.Obj.Decl).Elem().FieldByName("Values")
	if !objectValues.IsValid() {
		return result
	}

	values := objectValues.Interface().([]ast.Expr)

	for i := range values {
		switch item := values[i].(type) {
		case *ast.BasicLit:
			parsed := godmt.BasicTypeLiteralParser(d, item)
			result = append(result, parsed)

		case *ast.Ident:
			parsed := godmt.IdentifierParser(d, item)

			if parsed != nil {
				result = append(result, *parsed)
			}

		case *ast.CompositeLit:
			switch itemType := item.Type.(type) {
			case *ast.MapType:
				mapElements := reflect.ValueOf(item.Elts).Interface().([]ast.Expr)
				parsed := godmt.CompositeLiteralMapParser(d, mapElements, item)
				result = append(result, parsed)
			case *ast.ArrayType:
				sliceType := godmt.GetMapValueType(itemType.Elt)
				parsed := godmt.CompositeLiteralSliceParser(d, sliceType, item)
				result = append(result, parsed)
			}
		}
	}

	return result
}

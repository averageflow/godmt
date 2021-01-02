package godmt

import (
	"fmt"
	"go/ast"
	"reflect"
)

func ParseStruct(d *ast.Ident) []ScannedStruct {
	var result []ScannedStruct

	structTypes := reflect.ValueOf(d.Obj.Decl).Elem().FieldByName("Type")
	if !structTypes.IsValid() {
		return result
	}

	switch structTypes.Interface().(type) {
	case *ast.StructType:
		fields := structTypes.Interface().(*ast.StructType)
		fieldList := fields.Fields.List

		var rawScannedFields []ScannedStructField

		for i := range fieldList {
			parsed := ParseStructField(fieldList[i])
			if parsed != nil {
				rawScannedFields = append(rawScannedFields, *parsed)
			}
		}

		result = append(result, ScannedStruct{
			Doc:          nil,
			Name:         d.Name,
			Fields:       rawScannedFields,
			InternalType: StructType,
		})

	default:
		break
	}

	return result
}

func ParseStructField(item *ast.Field) *ScannedStructField {
	switch item.Type.(type) { //nolint:gocritic
	case *ast.Ident:
		return SimpleStructFieldParser(item)
	case *ast.StructType:
		return ParseNestedStruct(item)
	case *ast.SelectorExpr:
		return ImportedStructFieldParser(item)
	case *ast.MapType:
		return SimpleStructFieldParser(item)
	case *ast.ArrayType:
		return ParseComplexStructField(item.Names[0])
	case *ast.StarExpr:
		pointer := item.Type.(*ast.StarExpr).X
		switch value := pointer.(type) {
		case *ast.ArrayType:
			tag := item.Tag

			var tagValue string
			if tag != nil {
				tagValue = tag.Value
			}

			return &ScannedStructField{
				Doc:          ExtractComments(item.Doc),
				Name:         item.Names[0].Name,
				Kind:         GetSliceType(value),
				Tag:          tagValue,
				InternalType: SliceType,
				IsPointer:    true,
			}
		case *ast.Ident:
			tag := item.Tag

			var tagValue string
			if tag != nil {
				tagValue = tag.Value
			}
			return &ScannedStructField{
				Doc:          ExtractComments(item.Doc),
				Name:         item.Names[0].Name,
				Kind:         value.Name,
				Tag:          tagValue,
				InternalType: VarType,
				IsPointer:    true,
			}
		case *ast.MapType:
			key := GetMapValueType(value.Key)
			keyvalue := GetMapValueType(value.Value)
			kind := fmt.Sprintf("map[%s]%s", key, keyvalue)

			return &ScannedStructField{
				Doc:          ExtractComments(item.Doc),
				Name:         item.Names[0].Name,
				Kind:         kind,
				Tag:          item.Tag.Value,
				InternalType: MapType,
				IsPointer:    true,
			}
		default:
			return nil
		}
	default:
		return nil
	}
}

func ParseNestedStruct(field *ast.Field) *ScannedStructField {
	nestedFields := reflect.ValueOf(field.Type).Elem().FieldByName("Fields").Interface().(*ast.FieldList)

	var parsedNestedFields []ScannedStructField

	for i := range nestedFields.List {
		parsedField := ParseStructField(nestedFields.List[i])
		if parsedField != nil {
			parsedNestedFields = append(parsedNestedFields, *parsedField)
		}
	}

	tag := field.Tag

	var tagValue string
	if tag != nil {
		tagValue = tag.Value
	}

	return &ScannedStructField{
		Name:          field.Names[0].Name,
		Kind:          "struct",
		Tag:           tagValue,
		Doc:           ExtractComments(field.Doc),
		ImportDetails: nil,
		SubFields:     parsedNestedFields,
	}
}

func ParseComplexStructField(item *ast.Ident) *ScannedStructField {
	decl := item.Obj.Decl
	tag := reflect.ValueOf(decl).Elem().FieldByName("Tag").Interface().(*ast.BasicLit)
	comments := reflect.ValueOf(decl).Elem().FieldByName("Doc").Interface().(*ast.CommentGroup)

	objectType := reflect.ValueOf(decl).Elem().FieldByName("Type").Interface()

	var tagValue string
	if tag != nil {
		tagValue = tag.Value
	}

	result := &ScannedStructField{
		Name:          item.Name,
		Kind:          "",
		Tag:           tagValue,
		Doc:           ExtractComments(comments),
		ImportDetails: nil,
		InternalType:  0,
	}

	switch objectTypeDetails := objectType.(type) {
	case *ast.ArrayType:
		result.InternalType = SliceType

		switch sliceElement := objectTypeDetails.Elt.(type) {
		case *ast.StarExpr:
			result.Kind = fmt.Sprintf("[]%s", reflect.ValueOf(sliceElement.X).Elem().FieldByName("Name"))

		default:
			result.Kind = GetSliceType(objectTypeDetails)
		}

	default:
		return nil
	}

	return result
}

func ParseConstantsAndVariables(d *ast.Ident) []ScannedType {
	var result []ScannedType

	objectValues := reflect.ValueOf(d.Obj.Decl).Elem().FieldByName("Values")
	if !objectValues.IsValid() {
		return result
	}

	values := objectValues.Interface().([]ast.Expr)

	for i := range values {
		switch item := values[i].(type) {
		case *ast.BasicLit:
			parsed := BasicTypeLiteralParser(d, item)
			result = append(result, parsed)

		case *ast.Ident:
			parsed := IdentifierParser(d, item)

			if parsed != nil {
				result = append(result, *parsed)
			}

		case *ast.CompositeLit:
			switch itemType := item.Type.(type) {
			case *ast.MapType:
				mapElements := reflect.ValueOf(item.Elts).Interface().([]ast.Expr)
				parsed := CompositeLiteralMapParser(d, mapElements, item)
				result = append(result, parsed)
			case *ast.ArrayType:
				sliceType := GetMapValueType(itemType.Elt)
				parsed := CompositeLiteralSliceParser(d, sliceType, item)
				result = append(result, parsed)
			}
		}
	}

	return result
}

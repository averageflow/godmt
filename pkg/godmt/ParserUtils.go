package godmt

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
)

// BasicTypeLiteralParser will try to extract the type of a basic literal type,
// whether const or var.
func BasicTypeLiteralParser(d *ast.Ident, item *ast.BasicLit) ScannedType {
	itemType := fmt.Sprintf("%T", item.Value)

	if item.Kind == token.INT {
		itemType = "int64"
	} else if item.Kind == token.FLOAT {
		itemType = "float64"
	}

	rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)

	return ScannedType{
		Name:         d.Name,
		Kind:         itemType,
		Value:        item.Value,
		Doc:          ExtractComments(rawDecl.Doc),
		InternalType: ConstType,
	}
}

// IdentifierParser will try to extract the type of an identifier, like booleans.
// This will return a pointer to a ScannedType, thus the result should be checked for nil.
func IdentifierParser(d, item *ast.Ident) *ScannedType {
	rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)

	if item.Name == "true" || item.Name == "false" {
		return &ScannedType{
			Name:         d.Name,
			Kind:         "bool",
			Value:        item.Name,
			Doc:          ExtractComments(rawDecl.Doc),
			InternalType: ConstType,
		}
	}

	return nil
}

// CompositeLiteralSliceParser will extract a map ScannedType from the valid corresponding composite literal declaration.
func CompositeLiteralMapParser(d *ast.Ident, mapElements []ast.Expr, item *ast.CompositeLit) ScannedType {
	cleanMap := make(map[string]string)

	for j := range mapElements {
		rawKey := reflect.ValueOf(mapElements[j]).Elem().FieldByName("Key")
		switch rawKey.Interface().(type) {
		case *ast.BasicLit:
			rawValue := reflect.ValueOf(mapElements[j]).Elem().FieldByName("Value").Interface()
			switch item := rawValue.(type) {
			case *ast.BasicLit:
				cleanMap[fmt.Sprintf("%v", rawKey.Interface().(*ast.BasicLit).Value)] = item.Value
			case *ast.Ident:
				cleanMap[fmt.Sprintf("%v", rawKey.Interface().(*ast.BasicLit).Value)] = item.Name
			}
		default:
			break
		}
	}

	var doc []string

	rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)
	if rawDecl.Doc != nil {
		for i := range rawDecl.Doc.List {
			doc = append(doc, rawDecl.Doc.List[i].Text)
		}
	}

	return ScannedType{
		Name: d.Name,
		Kind: fmt.Sprintf(
			"map[%s]%s",
			GetMapValueType(item.Type.(*ast.MapType).Key),
			GetMapValueType(item.Type.(*ast.MapType).Value),
		),
		Value:        cleanMap,
		InternalType: MapType,
		Doc:          doc,
	}
}

// CompositeLiteralSliceParser will extract a slice ScannedType from the valid corresponding composite literal declaration.
func CompositeLiteralSliceParser(d *ast.Ident, sliceType string, item *ast.CompositeLit) ScannedType {
	rawDecl := reflect.ValueOf(d.Obj.Decl).Elem().Interface().(ast.ValueSpec)

	return ScannedType{
		Name:         d.Name,
		Kind:         fmt.Sprintf("[]%s", sliceType),
		Value:        ExtractSliceValues(item.Elts),
		Doc:          ExtractComments(rawDecl.Doc),
		InternalType: SliceType,
	}
}

// ImportedStructFieldParser will transform a field of a struct that contains a basic entity
// into a program readable ScannedStructField.
func SimpleStructFieldParser(field *ast.Field) *ScannedStructField {
	if field.Names != nil {
		// Basic type of a field inside a struct
		fieldType := reflect.ValueOf(field.Type).Elem().FieldByName("Name")

		tag := field.Tag

		var tagValue string
		if tag != nil {
			tagValue = tag.Value
		}

		return &ScannedStructField{
			Doc:  ExtractComments(field.Doc),
			Name: field.Names[0].Name,
			Kind: fieldType.Interface().(string),
			Tag:  tagValue,
		}
	}

	// Struct inside a struct
	fieldType := reflect.ValueOf(field.Type).Elem().FieldByName("Obj").Interface().(*ast.Object)
	tag := field.Tag

	var tagValue string
	if tag != nil {
		tagValue = tag.Value
	}

	return &ScannedStructField{
		Doc:  nil,
		Name: fieldType.Name,
		Kind: "struct",
		Tag:  tagValue,
	}
}

// ImportedStructFieldParser will transform a field of a struct that contains an imported entity
// into a program readable ScannedStructField.
func ImportedStructFieldParser(field *ast.Field) *ScannedStructField {
	fieldType := reflect.ValueOf(field.Type).Interface().(*ast.SelectorExpr)

	tag := field.Tag

	var tagValue string
	if tag != nil {
		tagValue = tag.Value
	}

	name := fmt.Sprintf("%s", field.Names)

	if len(field.Names) > 0 {
		name = field.Names[0].Name
	}

	packageName := fmt.Sprintf("%s", reflect.ValueOf(fieldType.X).Elem().FieldByName("Name"))

	return &ScannedStructField{
		Doc:  nil,
		Name: name,
		Kind: fieldType.Sel.Name,
		Tag:  tagValue,
		ImportDetails: &ImportedEntityDetails{
			EntityName:  fieldType.Sel.Name,
			PackageName: packageName,
		},
	}
}

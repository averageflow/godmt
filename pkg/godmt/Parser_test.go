package godmt

import (
	"fmt"
	"go/ast"
	"testing"

	"github.com/brianvoe/gofakeit/v5"
)

func TestParseStructRandomData(t *testing.T) {
	t.Parallel()

	var identity ast.Ident

	gofakeit.Struct(&identity)

	resultingStruct := ParseStruct(&identity)

	scannedType := fmt.Sprintf("%T", resultingStruct)

	if scannedType != "[]godmt.ScannedStruct" {
		t.Errorf("Expected []godmt.ScannedStruct got %s", scannedType)
	}
}

/*func TestParseStructExpectedType(t *testing.T) {
	var identity ast.Ident

	var fakeStruct ast.StructType

	gofakeit.Struct(&identity)
	gofakeit.Struct(&fakeStruct)

	identity.Obj.Decl = &fakeStruct

	v := reflect.ValueOf(identity.Obj.Decl).Elem().FieldByName("Type")
	if v.IsValid() {
		v.Set(reflect.ValueOf(fakeStruct))
	}

	resultingStruct := ParseStruct(&identity)

	fmt.Printf("%+v", resultingStruct)
}*/
func TestParseStructFieldIdent(t *testing.T) {
	t.Parallel()

	var field ast.Field

	gofakeit.Struct(&field)

	field.Type = ast.NewIdent(gofakeit.BeerAlcohol())
	scannedField := ParseStructField(&field)

	scannedType := fmt.Sprintf("%T", scannedField)

	if scannedType != "*godmt.ScannedStructField" {
		t.Errorf("Expected *godmt.ScannedStructField got %s", scannedType)
	}

	if scannedField.InternalType != ConstType {
		t.Errorf("Expected InternalType %d got %d", ConstType, scannedField.InternalType)
	}
}

package translators

import (
	"testing"

	"github.com/averageflow/godmt/pkg/godmt"
)

func TestIsEmbeddedStructForInheritance(t *testing.T) {
	sut := godmt.ScannedStructField{
		Name:          "",
		Kind:          "",
		Tag:           "",
		Doc:           nil,
		ImportDetails: nil,
		SubFields:     nil,
	}

	if IsEmbeddedStructForInheritance(&sut) {
		t.Errorf("Expected sut to not be an embedded struct for inheritance")
	}

	sut.Kind = StructTypeKeyWord
	sut.Tag = ""

	if !IsEmbeddedStructForInheritance(&sut) {
		t.Errorf("Expected sut to be an embedded struct for inheritance")
	}
}

func TestTransformTypeScriptRecord(t *testing.T) {
	testTable := map[string]string{
		"map[string]string":           "Record<string, string>",
		"map[string]int":              "Record<string, number>",
		"map[int]string":              "Record<number, string>",
		"map[string]interface{}":      "Record<string, any>",
		"map[string]map[int]string":   "Record<string, Record<number, string>>",
		"map[string][]map[int]string": "Record<string, Record<number, string>[]>",
		"map[int][]map[string]int":    "Record<number, Record<string, number>[]>",
		"map[string][]int":            "Record<string, number[]>",
	}

	for i := range testTable {
		sut := TransformTypeScriptRecord(i, false)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestTransformTypeScriptRecordPointer(t *testing.T) {
	testTable := map[string]string{
		"map[string]string":           "Record<string, string> | null",
		"map[string]int":              "Record<string, number> | null",
		"map[int]string":              "Record<number, string> | null",
		"map[string]interface{}":      "Record<string, any> | null",
		"map[string]map[int]string":   "Record<string, Record<number, string>> | null",
		"map[string][]map[int]string": "Record<string, Record<number, string>[]> | null",
		"map[int][]map[string]int":    "Record<number, Record<string, number>[]> | null",
		"map[string][]int":            "Record<string, number[]> | null",
	}

	for i := range testTable {
		sut := TransformTypeScriptRecord(i, true)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestTransformSwiftRecord(t *testing.T) {
	testTable := map[string]string{
		"map[string]string":           "Dictionary<String, String>",
		"map[string]int":              "Dictionary<String, Int>",
		"map[int]string":              "Dictionary<Int, String>",
		"map[string]interface{}":      "Dictionary<String, Any>",
		"map[string]map[int]string":   "Dictionary<String, Dictionary<Int, String>>",
		"map[string][]map[int]string": "Dictionary<String, [Dictionary<Int, String>]>",
		"map[int][]map[string]int":    "Dictionary<Int, [Dictionary<String, Int>]>",
		"map[string][]int":            "Dictionary<String, [Int]>",
	}

	for i := range testTable {
		sut := TransformSwiftRecord(i, false)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestTransformSwiftRecordPointer(t *testing.T) {
	testTable := map[string]string{
		"map[string]string":           "Dictionary<String, String>?",
		"map[string]int":              "Dictionary<String, Int>?",
		"map[int]string":              "Dictionary<Int, String>?",
		"map[string]interface{}":      "Dictionary<String, Any>?",
		"map[string]map[int]string":   "Dictionary<String, Dictionary<Int, String>>?",
		"map[string][]map[int]string": "Dictionary<String, [Dictionary<Int, String>]>?",
		"map[int][]map[string]int":    "Dictionary<Int, [Dictionary<String, Int>]>?",
		"map[string][]int":            "Dictionary<String, [Int]>?",
	}

	for i := range testTable {
		sut := TransformSwiftRecord(i, true)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestGetTypescriptCompatibleType(t *testing.T) {
	testTable := map[string]string{
		"int":         "number",
		"int32":       "number",
		"int64":       "number",
		"float":       "number",
		"float32":     "number",
		"float64":     "number",
		"string":      "string",
		"bool":        "boolean",
		"interface{}": "any",
		"NullFloat64": "(number | null)",
		"NullFloat32": "(number | null)",
		"NullInt32":   "(number | null)",
		"NullInt64":   "(number | null)",
		"NullString":  "(string | null)",
		"*int":        "(number | null)",
		"*string":     "(string | null)",
		"*float":      "(number | null)",
		"*bool":       "(boolean | null)",
	}

	for i := range testTable {
		sut := GetTypescriptCompatibleType(i, false)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestGetPHPCompatibleType(t *testing.T) {
	testTable := map[string]string{
		"int":         "int",
		"int32":       "int",
		"int64":       "int",
		"float":       "float",
		"float32":     "float",
		"float64":     "float",
		"string":      "string",
		"bool":        "bool",
		"interface{}": "",
		"NullFloat64": "?float",
		"NullFloat32": "?float",
		"NullInt32":   "?int",
		"NullInt64":   "?int",
		"NullString":  "?string",
		"*int":        "?int",
		"*string":     "?string",
		"*float":      "?float",
		"*bool":       "?bool",
	}

	for i := range testTable {
		sut := GetPHPCompatibleType(i, false)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestGetSwiftCompatibleType(t *testing.T) {
	testTable := map[string]string{
		"int":         "Int",
		"int32":       "Int",
		"int64":       "Int",
		"float":       "Float",
		"float32":     "Float",
		"float64":     "Float",
		"string":      "String",
		"bool":        "Bool",
		"interface{}": "Any",
		"NullFloat64": "Float?",
		"NullFloat32": "Float?",
		"NullInt32":   "Int?",
		"NullInt64":   "Int?",
		"NullString":  "String?",
		"*int":        "Int?",
		"*string":     "String?",
		"*float":      "Float?",
		"*bool":       "Bool?",
	}

	for i := range testTable {
		sut := GetSwiftCompatibleType(i, false)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

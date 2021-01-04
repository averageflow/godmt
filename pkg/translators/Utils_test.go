package translators

import (
	"testing"

	"github.com/averageflow/godmt/pkg/godmt"
)

func TestIsEmbeddedStructForInheritance(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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

func TestMapValuesToTypeScriptRecord(t *testing.T) {
	t.Parallel()

	testTable := map[string]string{
		"test1": "example1",
		"test2": "example2",
	}

	possibility1 := "\ttest1: example1,\n\ttest2: example2"
	possibility2 := "\ttest2: example2,\n\ttest1: example1"
	sut := MapValuesToTypeScriptRecord(testTable)
	if sut != possibility1 && sut != possibility2 {
		t.Errorf("Expected %s or %s, got %s", possibility1, possibility2, sut)
	}

}

func TestMapValuesToPHPArray(t *testing.T) {
	t.Parallel()

	testTable := map[string]string{
		"test1": "example1",
		"test2": "example2",
	}

	possibility1 := "\ttest1 => example1,\n\ttest2 => example2"
	possibility2 := "\ttest2 => example2,\n\ttest1 => example1"
	sut := MapValuesToPHPArray(testTable)
	if sut != possibility1 && sut != possibility2 {
		t.Errorf("Expected %s or %s, got %s", possibility1, possibility2, sut)
	}
}

func TestTransformSliceTypeToTypeScript(t *testing.T) {
	t.Parallel()

	testTable := map[string]string{
		"[]int":         "number[]",
		"[]float":       "number[]",
		"[]string":      "string[]",
		"[]bool":        "boolean[]",
		"[]interface{}": "any[]",
		"[]*int":        "(number | null)[]",
		"[]*string":     "(string | null)[]",
		"[]*float":      "(number | null)[]",
		"[]*bool":       "(boolean | null)[]",
	}

	for i := range testTable {
		sut := TransformSliceTypeToTypeScript(i, false)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestTransformSliceTypeToPHP(t *testing.T) {
	t.Parallel()

	testTable := map[string]string{
		"[]int":         "int[]",
		"[]float":       "float[]",
		"[]string":      "string[]",
		"[]bool":        "bool[]",
		"[]interface{}": "array",
		"[]*int":        "?int[]",
		"[]*string":     "?string[]",
		"[]*float":      "?float[]",
		"[]*bool":       "?bool[]",
	}

	for i := range testTable {
		sut := TransformSliceTypeToPHP(i, false)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestTransformSliceTypeToSwift(t *testing.T) {
	t.Parallel()

	testTable := map[string]string{
		"int":         "[Int]",
		"float":       "[Float]",
		"string":      "[String]",
		"bool":        "[Bool]",
		"interface{}": "[Any]",
		"*int":        "[Int?]",
		"*string":     "[String?]",
		"*float":      "[Float?]",
		"*bool":       "[Bool?]",
	}

	for i := range testTable {
		sut := TransformSliceTypeToSwift(i, false)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestTransformSliceTypeToTypeScriptPointer(t *testing.T) {
	t.Parallel()

	testTable := map[string]string{
		"[]int":         "number[] | null",
		"[]float":       "number[] | null",
		"[]string":      "string[] | null",
		"[]bool":        "boolean[] | null",
		"[]interface{}": "any[] | null",
		"[]*int":        "(number | null)[] | null",
		"[]*string":     "(string | null)[] | null",
		"[]*float":      "(number | null)[] | null",
		"[]*bool":       "(boolean | null)[] | null",
	}

	for i := range testTable {
		sut := TransformSliceTypeToTypeScript(i, true)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestTransformSliceTypeToPHPPointer(t *testing.T) {
	t.Parallel()

	testTable := map[string]string{
		"[]int":         "??int[]",
		"[]float":       "??float[]",
		"[]string":      "??string[]",
		"[]bool":        "??bool[]",
		"[]interface{}": "?array",
		"[]*int":        "??int[]",
		"[]*string":     "??string[]",
		"[]*float":      "??float[]",
		"[]*bool":       "??bool[]",
	}

	for i := range testTable {
		sut := TransformSliceTypeToPHP(i, true)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

func TestTransformSliceTypeToSwiftPointer(t *testing.T) {
	t.Parallel()

	testTable := map[string]string{
		"int":         "[Int]?",
		"float":       "[Float]?",
		"string":      "[String]?",
		"bool":        "[Bool]?",
		"interface{}": "[Any]?",
		"*int":        "[Int?]?",
		"*string":     "[String?]?",
		"*float":      "[Float?]?",
		"*bool":       "[Bool?]?",
	}

	for i := range testTable {
		sut := TransformSliceTypeToSwift(i, true)
		if sut != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], sut)
		}
	}
}

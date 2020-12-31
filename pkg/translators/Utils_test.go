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

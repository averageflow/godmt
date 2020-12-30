package translators

import (
	"fmt"
	"strings"

	"github.com/averageflow/godmt/pkg/godmt"

	"github.com/iancoleman/strcase"
)

func IsEmbeddedStructForInheritance(field *godmt.ScannedStructField) bool {
	return field.Kind == "struct" && field.Tag == ""
}

func GetTypescriptCompatibleType(goType string) string {
	result, ok := goTypeScriptTypeMappings[goType]
	if !ok {
		return goType
	}

	return result
}

func GetPHPCompatibleType(goType string) string {
	result, ok := goPHPTypeMappings[goType]
	if !ok {
		return goType
	}

	return result
}

func GetSwiftCompatibleType(goType string) string {
	result, ok := goSwiftTypeMappings[goType]
	if !ok {
		return goType
	}

	return result
}

func GetRecordType(goMapType string) string {
	var result string

	result = strings.ReplaceAll(goMapType, "map[", "")
	recordTypes := strings.Split(result, "]")

	for i := range recordTypes {
		recordTypes[i] = GetTypescriptCompatibleType(recordTypes[i])
	}

	result = strings.Join(recordTypes, ", ")

	return fmt.Sprintf("Record<%s>", result)
}

func GetDictionaryType(goMapType string) string {
	var result string

	result = strings.ReplaceAll(goMapType, "map[", "")
	recordTypes := strings.Split(result, "]")

	for i := range recordTypes {
		recordTypes[i] = GetSwiftCompatibleType(recordTypes[i])
	}

	result = strings.Join(recordTypes, ", ")

	return fmt.Sprintf("Dictionary<%s>", result)
}

func MapValuesToTypeScriptRecord(rawMap map[string]string) string {
	var entries []string //nolint:prealloc
	for i := range rawMap {
		entries = append(entries, fmt.Sprintf("\t%s: %s", i, rawMap[i]))
	}

	return strings.Join(entries, ",\n")
}

func MapValuesToPHPArray(rawMap map[string]string) string {
	var entries []string //nolint:prealloc
	for i := range rawMap {
		entries = append(entries, fmt.Sprintf("\t%s => %s", i, rawMap[i]))
	}

	return strings.Join(entries, ",\n")
}

func TransformSliceTypeToTypeScript(rawSliceType string) string {
	result := strings.ReplaceAll(rawSliceType, "[]", "")
	return fmt.Sprintf("%s[]", GetTypescriptCompatibleType(result))
}

func TransformSliceTypeToPHP(rawSliceType string) string {
	result := strings.ReplaceAll(rawSliceType, "[]", "")
	if result == "interface{}" {
		return "array"
	}
	return fmt.Sprintf("%s[]", GetPHPCompatibleType(result))
}

func TransformSliceTypeToSwift(rawSliceType string) string {
	result := strings.ReplaceAll(rawSliceType, "[]", "")
	return fmt.Sprintf("[%s]", GetSwiftCompatibleType(result))
}

func toCamelCase(raw string) string {
	return strcase.ToLowerCamel(raw)
}

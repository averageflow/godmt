package translators

import (
	"fmt"
	"regexp"
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

func GetRecordType(language string, goMapType string) string {
	var recordType string

	if language == SwiftTranslationMode {
		recordType = "Dictionary"
	} else if language == TypeScriptTranslationMode {
		recordType = "Record"
	}

	var re = regexp.MustCompile(`(?m)map\[[^]^[]+]`)

	var types []string

	for _, match := range re.FindAllString(goMapType, -1) {
		cleanMatch := strings.ReplaceAll(match, "map[", "")
		cleanMatch = strings.ReplaceAll(cleanMatch, "]", "")
		if language == SwiftTranslationMode {
			types = append(types, GetSwiftCompatibleType(cleanMatch))
		} else if language == TypeScriptTranslationMode {
			types = append(types, GetTypescriptCompatibleType(cleanMatch))
		}
	}

	typeSplit := strings.Split(goMapType, "]")
	lastType := typeSplit[len(typeSplit)-1]
	if language == SwiftTranslationMode {
		lastType = GetSwiftCompatibleType(lastType)
	} else if language == TypeScriptTranslationMode {
		lastType = GetTypescriptCompatibleType(lastType)
	}

	switch len(types) {
	case 3:
		return fmt.Sprintf("%s<%s, %s<%s, %s<%s, %s>>>", recordType, types[0], recordType, types[1], recordType, types[2], lastType)
	case 2:
		return fmt.Sprintf("%s<%s, %s<%s, %s>>", recordType, types[0], recordType, types[1], lastType)
	default:
		return fmt.Sprintf("%s<%s, %s>", recordType, types[0], lastType)
	}
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
	if strings.Contains(result, "map") {
		return fmt.Sprintf("%s[]", GetRecordType(TypeScriptTranslationMode, result))
	}
	return fmt.Sprintf("%s[]", GetTypescriptCompatibleType(result))
}

func TransformSliceTypeToPHP(rawSliceType string) string {
	result := strings.ReplaceAll(rawSliceType, "[]", "")
	if result == "interface{}" || strings.Contains(result, "map") {
		return "array"
	}

	return fmt.Sprintf("%s[]", GetPHPCompatibleType(result))
}

func TransformSliceTypeToSwift(rawSliceType string) string {
	result := strings.ReplaceAll(rawSliceType, "[]", "")
	if strings.Contains(result, "map") {
		return fmt.Sprintf("[%s]", GetRecordType(SwiftTranslationMode, result))
	}
	return fmt.Sprintf("[%s]", GetSwiftCompatibleType(result))
}

func toCamelCase(raw string) string {
	return strcase.ToLowerCamel(raw)
}

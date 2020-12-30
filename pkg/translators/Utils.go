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

type ScannedRecordItem struct {
	Kind       string
	IsMapSlice bool
}

func TransformTypeScriptRecord(goMapType string) string {
	var re = regexp.MustCompile(`(?m)\[?]?map\[[^]^[]+]`)

	var types []ScannedRecordItem
	lastType := ScannedRecordItem{
		Kind:       goMapType,
		IsMapSlice: false,
	}

	for _, match := range re.FindAllString(goMapType, -1) {
		if strings.HasPrefix(match, "[]map") {
			lastType.Kind = strings.ReplaceAll(lastType.Kind, match, "")
			cleanMatch := strings.ReplaceAll(match, "[]map[", "")
			cleanMatch = strings.ReplaceAll(cleanMatch, "]", "")
			types = append(types, ScannedRecordItem{Kind: GetTypescriptCompatibleType(cleanMatch), IsMapSlice: true})
		} else {
			lastType.Kind = strings.ReplaceAll(lastType.Kind, match, "")
			cleanMatch := strings.ReplaceAll(match, "map[", "")
			cleanMatch = strings.ReplaceAll(cleanMatch, "]", "")
			types = append(types, ScannedRecordItem{Kind: GetTypescriptCompatibleType(cleanMatch)})
		}
	}

	if strings.HasPrefix(lastType.Kind, "[]") {
		lastType.Kind = TransformSliceTypeToTypeScript(lastType.Kind)
	} else {

		lastType.Kind = GetTypescriptCompatibleType(lastType.Kind)
	}

	var result string

	for i := range types {
		if i != 0 && i <= len(types)-1 {
			result += ", "
		}

		result += "Record<"
		result += types[i].Kind
	}

	result += ", " + lastType.Kind
	for i := range types {
		if types[i].IsMapSlice {
			result += "[]>"
		} else {
			result += ">"
		}
	}

	return result
}

func TransformSwiftRecord(goMapType string) string {
	var re = regexp.MustCompile(`(?m)\[?]?map\[[^]^[]+]`)

	var types []ScannedRecordItem
	lastType := ScannedRecordItem{
		Kind:       goMapType,
		IsMapSlice: false,
	}

	for _, match := range re.FindAllString(goMapType, -1) {
		if strings.HasPrefix(match, "[]map") {
			lastType.Kind = strings.ReplaceAll(lastType.Kind, match, "")
			cleanMatch := strings.ReplaceAll(match, "[]map[", "")
			cleanMatch = strings.ReplaceAll(cleanMatch, "]", "")
			types = append(types, ScannedRecordItem{Kind: GetSwiftCompatibleType(cleanMatch), IsMapSlice: true})
		} else {
			lastType.Kind = strings.ReplaceAll(lastType.Kind, match, "")
			cleanMatch := strings.ReplaceAll(match, "map[", "")
			cleanMatch = strings.ReplaceAll(cleanMatch, "]", "")
			types = append(types, ScannedRecordItem{Kind: GetSwiftCompatibleType(cleanMatch)})
		}
	}

	if strings.HasPrefix(lastType.Kind, "[]") {
		lastType.Kind = TransformSliceTypeToSwift(lastType.Kind)
	} else {
		lastType.Kind = GetSwiftCompatibleType(lastType.Kind)
	}

	var result string

	for i := range types {
		if i != 0 && i <= len(types)-1 {
			result += ", "
		}

		if types[i].IsMapSlice {
			result += "["
		}
		result += "Dictionary<"
		result += types[i].Kind

	}

	result += ", " + lastType.Kind
	for i := range types {
		if types[i].IsMapSlice {
			result += "]>"
		} else {
			result += ">"
		}
	}

	return result
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
		return fmt.Sprintf("%s[]", TransformTypeScriptRecord(result))
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
		return fmt.Sprintf("[%s]", TransformSwiftRecord(result))
	}
	return fmt.Sprintf("[%s]", GetSwiftCompatibleType(result))
}

func toCamelCase(raw string) string {
	return strcase.ToLowerCamel(raw)
}

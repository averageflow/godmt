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

func GetTypescriptCompatibleType(goType string, isPointer bool) string {
	if strings.HasPrefix(goType, "*") {
		wantedType := strings.ReplaceAll(goType, "*", "")

		result, ok := goTypeScriptTypeMappings[wantedType]
		if !ok {
			return goType
		}

		return fmt.Sprintf("(%s | null)", result)
	}

	result, ok := goTypeScriptTypeMappings[goType]
	if !ok {
		return goType
	}

	if isPointer {
		return fmt.Sprintf("%s | null", result)
	}

	return result
}

func GetPHPCompatibleType(goType string, isPointer bool) string {
	if strings.HasPrefix(goType, "*") {
		wantedType := strings.ReplaceAll(goType, "*", "")

		result, ok := goPHPTypeMappings[wantedType]
		if !ok {
			return goType
		}

		return fmt.Sprintf("?%s", result)
	}

	result, ok := goPHPTypeMappings[goType]
	if !ok {
		return goType
	}

	if isPointer {
		return fmt.Sprintf("?%s", result)
	}

	return result
}

func GetSwiftCompatibleType(goType string, isPointer bool) string {
	if strings.HasPrefix(goType, "*") {
		wantedType := strings.ReplaceAll(goType, "*", "")

		result, ok := goSwiftTypeMappings[wantedType]
		if !ok {
			return goType
		}

		return fmt.Sprintf("%s?", result)
	}

	result, ok := goSwiftTypeMappings[goType]
	if !ok {
		return goType
	}

	if isPointer {
		return fmt.Sprintf("%s?", result)
	}

	return result
}

type ScannedRecordItem struct {
	Kind       string
	IsMapSlice bool
}

func TransformTypeScriptRecord(goMapType string, isPointer bool) string {
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

			types = append(types, ScannedRecordItem{Kind: GetTypescriptCompatibleType(cleanMatch, false), IsMapSlice: true})
		} else {
			lastType.Kind = strings.ReplaceAll(lastType.Kind, match, "")
			cleanMatch := strings.ReplaceAll(match, "map[", "")
			cleanMatch = strings.ReplaceAll(cleanMatch, "]", "")

			types = append(types, ScannedRecordItem{Kind: GetTypescriptCompatibleType(cleanMatch, false)})
		}
	}

	if strings.HasPrefix(lastType.Kind, "[]") {
		lastType.Kind = TransformSliceTypeToTypeScript(lastType.Kind, false)
	} else if strings.HasPrefix(lastType.Kind, "*[]") {
		kind := strings.ReplaceAll(lastType.Kind, "*", "")
		lastType.Kind = TransformSliceTypeToTypeScript(kind, true)
	} else {
		lastType.Kind = GetTypescriptCompatibleType(lastType.Kind, false)
	}

	var result string

	for i := range types {
		if i != 0 && i <= len(types)-1 {
			result += ", " //nolint:goconst
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

	if isPointer {
		return fmt.Sprintf("%s | null", result)
	}

	return result
}

func TransformSwiftRecord(goMapType string, isPointer bool) string {
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

			types = append(types, ScannedRecordItem{Kind: GetSwiftCompatibleType(cleanMatch, false), IsMapSlice: true})
		} else {
			lastType.Kind = strings.ReplaceAll(lastType.Kind, match, "")
			cleanMatch := strings.ReplaceAll(match, "map[", "")
			cleanMatch = strings.ReplaceAll(cleanMatch, "]", "")
			types = append(types, ScannedRecordItem{Kind: GetSwiftCompatibleType(cleanMatch, false)})
		}
	}

	if strings.HasPrefix(lastType.Kind, "[]") {
		lastType.Kind = TransformSliceTypeToSwift(lastType.Kind, false)
	} else if strings.HasPrefix(lastType.Kind, "*[]") {
		kind := strings.ReplaceAll(lastType.Kind, "*", "")
		lastType.Kind = TransformSliceTypeToSwift(kind, true)

	} else {
		lastType.Kind = GetSwiftCompatibleType(lastType.Kind, false)
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

	if isPointer {
		return fmt.Sprintf("%s?", result)
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

func TransformSliceTypeToTypeScript(rawSliceType string, isPointer bool) string {
	result := strings.ReplaceAll(rawSliceType, "[]", "")
	if strings.Contains(result, "map") {
		result = fmt.Sprintf("%s[]", TransformTypeScriptRecord(result, false))
	} else {
		result = fmt.Sprintf("%s[]", GetTypescriptCompatibleType(result, false))
	}

	if isPointer {
		return fmt.Sprintf("%s | null", result)
	}

	return result
}

func TransformSliceTypeToPHP(rawSliceType string, isPointer bool) string {
	result := strings.ReplaceAll(rawSliceType, "[]", "")
	if result == "interface{}" || strings.Contains(result, "map") {
		result = "array"
	} else {
		result = fmt.Sprintf("%s[]", GetPHPCompatibleType(result, isPointer))
	}

	if isPointer {
		return fmt.Sprintf("?%s", result)
	}

	return result
}

func TransformSliceTypeToSwift(rawSliceType string, isPointer bool) string {
	result := strings.ReplaceAll(rawSliceType, "[]", "")
	if strings.Contains(result, "map") {
		result = fmt.Sprintf("[%s]", TransformSwiftRecord(result, false))
	} else {
		result = fmt.Sprintf("[%s]", GetSwiftCompatibleType(result, false))
	}

	if isPointer {
		return fmt.Sprintf("%s?", result)
	}

	return result
}

func toCamelCase(raw string) string {
	return strcase.ToLowerCamel(raw)
}

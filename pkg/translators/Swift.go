package translators

import (
	"fmt"

	"github.com/averageflow/godmt/pkg/godmt"
)

var goSwiftTypeMappings = map[string]string{ //nolint:gochecknoglobals
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
}

type SwiftTranslator struct {
	Translator
}

func (t *SwiftTranslator) Translate() string { //nolint:gocognit,gocyclo
	var result string

	for i := range t.Data.ConstantSort { //nolint:dupl
		entity := t.Data.ScanResult[t.Data.ConstantSort[i]]
		if len(entity.Doc) > 0 {
			for j := range entity.Doc {
				result += fmt.Sprintf("%s\n", entity.Doc[j])
			}
		}

		switch entity.InternalType {
		case godmt.ConstType:
			result += fmt.Sprintf(
				"let %s: %s = %s\n\n",
				entity.Name,
				GetSwiftCompatibleType(entity.Kind),
				entity.Value,
			)
		case godmt.MapType:
			result += fmt.Sprintf(
				"let %s: %s = [\n",
				entity.Name,
				TransformSwiftRecord(entity.Kind),
			)
			result += fmt.Sprintf("%s\n", MapValuesToTypeScriptRecord(entity.Value.(map[string]string)))
			result += "]\n\n" //nolint:goconst
		case godmt.SliceType:
			result += fmt.Sprintf(
				"var %s: %s = [\n",
				entity.Name,
				TransformSliceTypeToSwift(entity.Kind),
			)
			result += fmt.Sprintf("%s\n", godmt.SliceValuesToPrettyList(entity.Value.([]string)))
			result += "];\n\n" //nolint:goconst
		}
	}

	for i := range t.Data.StructSort {
		var inheritedFields []godmt.ScannedStructField

		entity := t.Data.StructScanResult[t.Data.StructSort[i]]

		for j := range entity.Fields {
			if IsEmbeddedStructForInheritance(&entity.Fields[j]) {
				inheritedFields = append(inheritedFields, t.Data.StructScanResult[entity.Fields[j].Name].Fields...)
			}
		}

		var structExtension string

		structExtension += "\n\tenum CodingKeys: String, CodingKey {\n"

		result += fmt.Sprintf("\nstruct %s: Decodable {\n", entity.Name)

		mergedWithInheritedFields := entity.Fields
		mergedWithInheritedFields = append(mergedWithInheritedFields, inheritedFields...)

		for j := range mergedWithInheritedFields {
			entityField := mergedWithInheritedFields[j]
			if IsEmbeddedStructForInheritance(&entityField) {
				continue
			}

			tag := godmt.CleanTagName(entityField.Tag)
			if tag == "" || t.Preserve {
				tag = entityField.Name
			}

			if entityField.Doc != nil {
				for k := range entityField.Doc {
					result += fmt.Sprintf("\t%s\n", entityField.Doc[k])
				}
			}

			if len(entityField.SubFields) > 0 {
				result += fmt.Sprintf("\tstruct %s {\n", quoteWhenNeeded(tag))

				for k := range entityField.SubFields {
					subtag := godmt.CleanTagName(entityField.SubFields[k].Tag)
					if subtag == "" || t.Preserve {
						subtag = entityField.SubFields[k].Name
					}

					result += fmt.Sprintf("\t\tvar %s: %s;\n", quoteWhenNeeded(subtag), GetSwiftCompatibleType(entityField.SubFields[k].Kind))
				}

				result += "\t}\n" //nolint:goconst
			} else {
				switch entityField.InternalType {
				case godmt.MapType:
					result += fmt.Sprintf("\tvar %s: %s\n", tag, TransformSwiftRecord(entityField.Kind))
				case godmt.SliceType:
					result += fmt.Sprintf("\tvar %s: %s\n", tag, TransformSliceTypeToSwift(entityField.Kind))
				default:
					result += fmt.Sprintf("\tvar %s: %s\n", toCamelCase(tag), GetSwiftCompatibleType(entityField.Kind))
				}
			}

			structExtension += fmt.Sprintf("\t\tcase %s = \"%s\"\n", toCamelCase(tag), tag)
		}

		structExtension += "\t}\n" //nolint:goconst
		result += structExtension
		result += "}\n" //nolint:goconst
	}

	return result
}

package translators

import (
	"fmt"
	"strings"

	"github.com/averageflow/godmt/pkg/godmt"
)

var goTypeScriptTypeMappings = map[string]string{ //nolint:gochecknoglobals
	"int":         "number",
	"int32":       "number",
	"int64":       "number",
	"float":       "number",
	"float32":     "number",
	"float64":     "number",
	"string":      "string",
	"bool":        "boolean",
	"interface{}": "any",
	"NullFloat64": "number | null",
	"NullFloat32": "number | null",
	"NullInt32":   "number | null",
	"NullInt64":   "number | null",
	"NullString":  "string | null",
}

type TypeScriptTranslator struct {
	Translator
}

func (t *TypeScriptTranslator) Translate() string {
	var imports string

	var result string

	for i := range t.Data.ConstantSort {
		entity := t.Data.ScanResult[t.Data.ConstantSort[i]]
		if len(entity.Doc) > 0 {
			for j := range entity.Doc {
				result += fmt.Sprintf("%s\n", entity.Doc[j])
			}
		}

		switch entity.InternalType {
		case godmt.ConstType:
			result += fmt.Sprintf(
				"export const %s: %s = %s;\n\n",
				entity.Name,
				GetTypescriptCompatibleType(entity.Kind),
				entity.Value,
			)
		case godmt.MapType:
			result += fmt.Sprintf(
				"export const %s: %s = {\n",
				entity.Name,
				GetRecordType(entity.Kind),
			)
			result += fmt.Sprintf("%s\n", MapValuesToTypeScriptRecord(entity.Value.(map[string]string)))
			result += "};\n\n"
		case godmt.SliceType:
			result += fmt.Sprintf(
				"export const %s: %s = [\n",
				entity.Name,
				TransformSliceTypeToTypeScript(entity.Kind),
			)
			result += fmt.Sprintf("%s\n", godmt.SliceValuesToPrettyList(entity.Value.([]string)))
			result += "];\n\n"
		}
	}

	for i := range t.Data.StructSort {
		var extendsClasses []string

		entity := t.Data.StructScanResult[t.Data.StructSort[i]]
		for j := range entity.Fields {
			if IsEmbeddedStructForInheritance(&entity.Fields[j]) {
				extendsClasses = append(extendsClasses, entity.Fields[j].Name)
			}
		}

		result += fmt.Sprintf("\nexport interface %s", entity.Name)
		if len(extendsClasses) > 0 {
			result += fmt.Sprintf(" extends %s", strings.Join(extendsClasses, ", "))
		}

		result += " {\n"

		for j := range entity.Fields {
			entityField := entity.Fields[j]
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
				result += fmt.Sprintf("\t%s: {\n", quoteWhenNeeded(tag))
				for k := range entityField.SubFields {
					subtag := godmt.CleanTagName(entityField.SubFields[k].Tag)
					if subtag == "" || t.Preserve {
						subtag = entityField.SubFields[k].Name
					}
					result += fmt.Sprintf("\t\t%s: %s;\n", quoteWhenNeeded(subtag), GetTypescriptCompatibleType(entityField.SubFields[k].Kind))
				}
				result += "\t}\n"
			} else {
				result += fmt.Sprintf("\t%s: %s;\n", quoteWhenNeeded(tag), GetTypescriptCompatibleType(entityField.Kind))
			}

			if entityField.ImportDetails != nil {
				imports += fmt.Sprintf(
					"import { %s } from \"%s\";\n",
					entityField.ImportDetails.EntityName,
					entityField.ImportDetails.PackageName,
				)
			}
		}

		result += "}\n"
	}

	if imports != "" {
		return fmt.Sprintf("%s\n\n%s", imports, result)
	}

	return result
}

func quoteWhenNeeded(raw string) string {
	if strings.Contains(raw, ":") {
		return fmt.Sprintf(`"%s"`, raw)
	}

	return raw
}

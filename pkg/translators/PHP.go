package translators

import (
	"fmt"
	"strings"

	"github.com/averageflow/godmt/pkg/godmt"
)

var goPHPTypeMappings = map[string]string{ //nolint:gochecknoglobals
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
}

type PHPTranslator struct {
	Translator
}

func (t *PHPTranslator) Translate() string { //nolint:gocognit,gocyclo
	var imports string

	result := "<?php\n\n"

	for i := range t.Data.ConstantSort { //nolint:dupl
		entity := t.Data.ScanResult[t.Data.ConstantSort[i]]

		switch entity.InternalType {
		case godmt.ConstType:
			result += "/**\n"
			if len(entity.Doc) > 0 {
				for j := range entity.Doc {
					result += fmt.Sprintf(" * %s\n", strings.ReplaceAll(entity.Doc[j], "// ", ""))
				}
			}
			result += fmt.Sprintf(
				" * @const %s %s\n */\n",
				entity.Name,
				GetPHPCompatibleType(entity.Kind),
			)
			result += fmt.Sprintf(
				"const %s = %s;\n\n",
				entity.Name,
				entity.Value,
			)
		case godmt.MapType:
			result += "/**\n"
			if len(entity.Doc) > 0 {
				for j := range entity.Doc {
					result += fmt.Sprintf(" * %s\n", strings.ReplaceAll(entity.Doc[j], "// ", ""))
				}
			}
			result += fmt.Sprintf(" * @const array %s\n */\n", entity.Name)
			result += fmt.Sprintf(
				"const %s = [\n",
				entity.Name,
			)
			result += fmt.Sprintf("%s\n", MapValuesToPHPArray(entity.Value.(map[string]string)))
			result += "];\n\n"
		case godmt.SliceType:
			result += "/**\n"
			if len(entity.Doc) > 0 {
				for j := range entity.Doc {
					result += fmt.Sprintf(" * %s\n", strings.ReplaceAll(entity.Doc[j], "// ", ""))
				}
			}
			result += fmt.Sprintf(
				" * @const %s %s\n */\n",
				TransformSliceTypeToPHP(entity.Kind),
				entity.Name,
			)

			result += fmt.Sprintf(
				"const %s = [\n",
				entity.Name,
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

		result += fmt.Sprintf("\nclass %s", entity.Name)
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
				result += fmt.Sprintf("\tprotected %s $%s;\n", GetPHPCompatibleType(entityField.Kind), tag)
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

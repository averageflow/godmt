package translators

import (
	"fmt"
	"strings"

	"github.com/averageflow/goschemaconverter/pkg/syntaxtreeparser"
)

var goTypeScriptTypeMappings = map[string]string{
	"int":         "number",
	"int32":       "number",
	"int64":       "number",
	"float":       "number",
	"float32":     "number",
	"float64":     "number",
	"string":      "string",
	"bool":        "boolean",
	"interface{}": "any",
}

type TypeScriptTranslator struct {
	Translator
}

func (t *TypeScriptTranslator) Translate() string {
	fmt.Println(`
----------------------------------
Performing TypeScript translation!
----------------------------------
	`)

	var imports string
	var result string

	for i := range t.OrderedTypes {
		if len(t.ScannedTypes[t.OrderedTypes[i]].Doc) > 0 {
			for j := range t.ScannedTypes[t.OrderedTypes[i]].Doc {
				result += fmt.Sprintf("%s\n", t.ScannedTypes[t.OrderedTypes[i]].Doc[j])
			}
		}

		switch t.ScannedTypes[t.OrderedTypes[i]].InternalType {
		case syntaxtreeparser.ConstType:
			result += fmt.Sprintf(
				"export const %s: %s = %s;\n\n",
				t.ScannedTypes[t.OrderedTypes[i]].Name,
				getTypescriptCompatibleType(t.ScannedTypes[t.OrderedTypes[i]].Kind),
				t.ScannedTypes[t.OrderedTypes[i]].Value,
			)
		case syntaxtreeparser.MapType:
			result += fmt.Sprintf(
				"export const %s: %s = {\n",
				t.ScannedTypes[t.OrderedTypes[i]].Name,
				getRecordType(t.ScannedTypes[t.OrderedTypes[i]].Kind),
			)
			result += fmt.Sprintf("%s\n", mapValuesToTypeScriptRecord(t.ScannedTypes[t.OrderedTypes[i]].Value.(map[string]string)))
			result += fmt.Sprint("};\n\n")
		case syntaxtreeparser.SliceType:
			result += fmt.Sprintf(
				"export const %s: %s = [\n",
				t.ScannedTypes[t.OrderedTypes[i]].Name,
				transformSliceTypeToTypeScript(t.ScannedTypes[t.OrderedTypes[i]].Kind),
			)
			result += fmt.Sprintf("%s\n", syntaxtreeparser.SliceValuesToPrettyList(t.ScannedTypes[t.OrderedTypes[i]].Value.([]string)))

			result += fmt.Sprint("];\n\n")
		}

	}

	for i := range t.OrderedStructs {
		var extendsClasses []string
		for j := range t.ScannedStructs[t.OrderedStructs[i]].Fields {
			if isEmbeddedStructForInheritance(t.ScannedStructs[t.OrderedStructs[i]].Fields[j]) {
				extendsClasses = append(extendsClasses, t.ScannedStructs[t.OrderedStructs[i]].Fields[j].Name)
			}
		}

		result += fmt.Sprintf("\nexport interface %s", t.ScannedStructs[t.OrderedStructs[i]].Name)
		if len(extendsClasses) > 0 {
			result += fmt.Sprintf(" extends %s", strings.Join(extendsClasses, ", "))
		}

		result += fmt.Sprint(" {\n")

		for j := range t.ScannedStructs[t.OrderedStructs[i]].Fields {
			if isEmbeddedStructForInheritance(t.ScannedStructs[t.OrderedStructs[i]].Fields[j]) {
				continue
			}

			tag := CleanTagName(t.ScannedStructs[t.OrderedStructs[i]].Fields[j].Tag)
			if tag == "" || t.Preserve {
				tag = t.ScannedStructs[t.OrderedStructs[i]].Fields[j].Name
			}

			if t.ScannedStructs[t.OrderedStructs[i]].Fields[j].Doc != nil {
				for k := range t.ScannedStructs[t.OrderedStructs[i]].Fields[j].Doc {
					result += fmt.Sprintf("\t%s\n", t.ScannedStructs[t.OrderedStructs[i]].Fields[j].Doc[k])
				}
			}

			result += fmt.Sprintf("\t%s: %s;\n", tag, getTypescriptCompatibleType(t.ScannedStructs[t.OrderedStructs[i]].Fields[j].Kind))

			if t.ScannedStructs[t.OrderedStructs[i]].Fields[j].ImportDetails != nil {
				imports += fmt.Sprintf(
					"import { %s } from \"%s\";\n",
					t.ScannedStructs[t.OrderedStructs[i]].Fields[j].ImportDetails.EntityName,
					t.ScannedStructs[t.OrderedStructs[i]].Fields[j].ImportDetails.PackageName,
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

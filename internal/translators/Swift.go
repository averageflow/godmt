package translators

import (
	"fmt"

	"github.com/averageflow/goschemaconverter/pkg/syntaxtreeparser"
)

var goSwiftTypeMappings = map[string]string{
	"int":         "Int",
	"int32":       "Int",
	"int64":       "Int",
	"float":       "Float",
	"float32":     "Float",
	"float64":     "Float",
	"string":      "String",
	"bool":        "Bool",
	"interface{}": "Any",
}

type SwiftTranslator struct {
	Translator
}

func (t *SwiftTranslator) Translate() string {
	fmt.Println(`
----------------------------------
Performing Swift translation!
----------------------------------
	`)

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
				"let %s: %s = %s\n\n",
				t.ScannedTypes[t.OrderedTypes[i]].Name,
				getSwiftCompatibleType(t.ScannedTypes[t.OrderedTypes[i]].Kind),
				t.ScannedTypes[t.OrderedTypes[i]].Value,
			)
		case syntaxtreeparser.MapType:
			result += fmt.Sprintf(
				"let %s: %s = [\n",
				t.ScannedTypes[t.OrderedTypes[i]].Name,
				getDictionaryType(t.ScannedTypes[t.OrderedTypes[i]].Kind),
			)
			result += fmt.Sprintf("%s\n", mapValuesToTypeScriptRecord(t.ScannedTypes[t.OrderedTypes[i]].Value.(map[string]string)))
			result += fmt.Sprint("]\n\n")
		case syntaxtreeparser.SliceType:
			result += fmt.Sprintf(
				"var %s: %s = [\n",
				t.ScannedTypes[t.OrderedTypes[i]].Name,
				transformSliceTypeToSwift(t.ScannedTypes[t.OrderedTypes[i]].Kind),
			)
			result += fmt.Sprintf("%s\n", syntaxtreeparser.SliceValuesToPrettyList(t.ScannedTypes[t.OrderedTypes[i]].Value.([]string)))

			result += fmt.Sprint("];\n\n")
		}

	}

	for i := range t.OrderedStructs {
		var extendsClasses []string
		var inheritedFields []syntaxtreeparser.ScannedStructField

		for j := range t.ScannedStructs[t.OrderedStructs[i]].Fields {
			if isEmbeddedStructForInheritance(t.ScannedStructs[t.OrderedStructs[i]].Fields[j]) {
				extendsClasses = append(extendsClasses, t.ScannedStructs[t.OrderedStructs[i]].Fields[j].Name)
				inheritedFields = append(inheritedFields, t.ScannedStructs[t.ScannedStructs[t.OrderedStructs[i]].Fields[j].Name].Fields...)
			}
		}

		var structExtension string

		structExtension += "\n\tenum CodingKeys: String, CodingKey {\n"

		result += fmt.Sprintf("\nstruct %s: Decodable {\n", t.ScannedStructs[t.OrderedStructs[i]].Name)

		mergedWithInheritedFields := t.ScannedStructs[t.OrderedStructs[i]].Fields
		mergedWithInheritedFields = append(mergedWithInheritedFields, inheritedFields...)

		for j := range mergedWithInheritedFields {
			if isEmbeddedStructForInheritance(mergedWithInheritedFields[j]) {
				continue
			}

			tag := CleanTagName(mergedWithInheritedFields[j].Tag)
			if tag == "" || t.Preserve {
				tag = mergedWithInheritedFields[j].Name
			}

			if mergedWithInheritedFields[j].Doc != nil {
				for k := range mergedWithInheritedFields[j].Doc {
					result += fmt.Sprintf("\t%s\n", mergedWithInheritedFields[j].Doc[k])
				}
			}

			result += fmt.Sprintf("\tvar %s: %s\n", toCamelCase(tag), getSwiftCompatibleType(mergedWithInheritedFields[j].Kind))
			structExtension += fmt.Sprintf("\t\tcase %s = \"%s\"\n", toCamelCase(tag), tag)
		}

		structExtension += "\t}\n"
		result += structExtension

		result += "}\n"
	}

	return result
}

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
		entity := t.ScannedTypes[t.OrderedTypes[i]]
		if len(entity.Doc) > 0 {
			for j := range entity.Doc {
				result += fmt.Sprintf("%s\n", entity.Doc[j])
			}
		}

		switch entity.InternalType {
		case syntaxtreeparser.ConstType:
			result += fmt.Sprintf(
				"let %s: %s = %s\n\n",
				entity.Name,
				getSwiftCompatibleType(entity.Kind),
				entity.Value,
			)
		case syntaxtreeparser.MapType:
			result += fmt.Sprintf(
				"let %s: %s = [\n",
				entity.Name,
				getDictionaryType(entity.Kind),
			)
			result += fmt.Sprintf("%s\n", mapValuesToTypeScriptRecord(entity.Value.(map[string]string)))
			result += fmt.Sprint("]\n\n")
		case syntaxtreeparser.SliceType:
			result += fmt.Sprintf(
				"var %s: %s = [\n",
				entity.Name,
				transformSliceTypeToSwift(entity.Kind),
			)
			result += fmt.Sprintf("%s\n", syntaxtreeparser.SliceValuesToPrettyList(entity.Value.([]string)))
			result += fmt.Sprint("];\n\n")
		}
	}

	for i := range t.OrderedStructs {
		var extendsClasses []string
		var inheritedFields []syntaxtreeparser.ScannedStructField

		entity := t.ScannedStructs[t.OrderedStructs[i]]

		for j := range entity.Fields {
			if isEmbeddedStructForInheritance(entity.Fields[j]) {
				extendsClasses = append(extendsClasses, entity.Fields[j].Name)
				inheritedFields = append(inheritedFields, t.ScannedStructs[entity.Fields[j].Name].Fields...)
			}
		}

		var structExtension string

		structExtension += "\n\tenum CodingKeys: String, CodingKey {\n"

		result += fmt.Sprintf("\nstruct %s: Decodable {\n", entity.Name)

		mergedWithInheritedFields := entity.Fields
		mergedWithInheritedFields = append(mergedWithInheritedFields, inheritedFields...)

		for j := range mergedWithInheritedFields {
			entityField := mergedWithInheritedFields[j]
			if isEmbeddedStructForInheritance(entityField) {
				continue
			}

			tag := syntaxtreeparser.CleanTagName(entityField.Tag)
			if tag == "" || t.Preserve {
				tag = entityField.Name
			}

			if entityField.Doc != nil {
				for k := range entityField.Doc {
					result += fmt.Sprintf("\t%s\n", entityField.Doc[k])
				}
			}

			result += fmt.Sprintf("\tvar %s: %s\n", toCamelCase(tag), getSwiftCompatibleType(entityField.Kind))
			structExtension += fmt.Sprintf("\t\tcase %s = \"%s\"\n", toCamelCase(tag), tag)
		}

		structExtension += "\t}\n"
		result += structExtension
		result += "}\n"
	}

	return result
}

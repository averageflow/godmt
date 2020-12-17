package translators

import (
	"fmt"

	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
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
	preserve       bool
	scannedTypes   map[string]syntaxtree.ScannedType
	scannedStructs map[string]syntaxtree.ScannedStruct
}

func (t *SwiftTranslator) Setup(preserve bool, d map[string]syntaxtree.ScannedType, s map[string]syntaxtree.ScannedStruct) {
	t.preserve = preserve
	t.scannedTypes = d
	t.scannedStructs = s
}

func (t *SwiftTranslator) Translate() string {
	fmt.Println(`
----------------------------------
Performing Swift translation!
----------------------------------
	`)

	var imports string
	var result string

	for i := range t.scannedTypes {
		if len(t.scannedTypes[i].Doc) > 0 {
			for j := range t.scannedTypes[i].Doc {
				result += fmt.Sprintf("%s\n", t.scannedTypes[i].Doc[j])
			}
		}

		switch t.scannedTypes[i].InternalType {
		case syntaxtree.ConstType:
			result += fmt.Sprintf(
				"let %s: %s = %s\n\n",
				t.scannedTypes[i].Name,
				getSwiftCompatibleType(t.scannedTypes[i].Kind),
				t.scannedTypes[i].Value,
			)
		case syntaxtree.MapType:
			result += fmt.Sprintf(
				"let %s: %s = [\n",
				t.scannedTypes[i].Name,
				getDictionaryType(t.scannedTypes[i].Kind),
			)
			result += fmt.Sprintf("%s\n", mapValuesToTypeScriptRecord(t.scannedTypes[i].Value.(map[string]string)))
			result += fmt.Sprint("]\n\n")
		case syntaxtree.SliceType:
			result += fmt.Sprintf(
				"var %s: %s = [\n",
				t.scannedTypes[i].Name,
				transformSliceTypeToSwift(t.scannedTypes[i].Kind),
			)
			result += fmt.Sprintf("%s\n", sliceValuesToPrettyList(t.scannedTypes[i].Value.([]string)))

			result += fmt.Sprint("];\n\n")
		}

	}

	for i := range t.scannedStructs {
		var extendsClasses []string
		var inheritedFields []syntaxtree.ScannedStructField

		for j := range t.scannedStructs[i].Fields {
			if isEmbeddedStructForInheritance(t.scannedStructs[i].Fields[j]) {
				extendsClasses = append(extendsClasses, t.scannedStructs[i].Fields[j].Name)
				inheritedFields = append(inheritedFields, t.scannedStructs[t.scannedStructs[i].Fields[j].Name].Fields...)
			}
		}

		result += fmt.Sprintf("\nstruct %s: Decodable {\n", t.scannedStructs[i].Name)

		mergedWithInheritedFields := t.scannedStructs[i].Fields
		mergedWithInheritedFields = append(mergedWithInheritedFields, inheritedFields...)

		for j := range mergedWithInheritedFields {
			if isEmbeddedStructForInheritance(mergedWithInheritedFields[j]) {
				continue
			}

			tag := CleanTagName(mergedWithInheritedFields[j].Tag)
			if tag == "" || t.preserve {
				tag = mergedWithInheritedFields[j].Name
			}

			if mergedWithInheritedFields[j].Doc != nil {
				for k := range mergedWithInheritedFields[j].Doc {
					result += fmt.Sprintf("\t%s\n", mergedWithInheritedFields[j].Doc[k])
				}
			}

			result += fmt.Sprintf("\tvar %s: %s\n", tag, getSwiftCompatibleType(mergedWithInheritedFields[j].Kind))

			if mergedWithInheritedFields[j].ImportDetails != nil {
				imports += fmt.Sprintf(
					"import { %s } from \"%s\";\n",
					mergedWithInheritedFields[j].ImportDetails.EntityName,
					mergedWithInheritedFields[j].ImportDetails.PackageName,
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

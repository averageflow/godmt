package translators

import (
	"fmt"
	"strings"

	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
)

var goTypeScriptTypeMappings = map[string]string{
	"int":         "number",
	"int32":       "number",
	"int64":       "number",
	"float":       "number",
	"float32":     "number",
	"float64":     "number",
	"string":      "string",
	"bool":        "bool",
	"interface{}": "any",
}

type TypeScriptTranslator struct {
	preserve       bool
	scannedTypes   []syntaxtree.ScannedType
	scannedStructs []syntaxtree.ScannedStruct
}

func (t *TypeScriptTranslator) Setup(preserve bool, d []syntaxtree.ScannedType, s []syntaxtree.ScannedStruct) {
	t.preserve = preserve
	t.scannedTypes = d
	t.scannedStructs = s
}

func (t *TypeScriptTranslator) Translate() string {
	fmt.Println(`
----------------------------------
Performing TypeScript translation!
----------------------------------
	`)

	var result string

	for i := range t.scannedTypes {
		if len(t.scannedTypes[i].Doc) > 0 {
			for j := range t.scannedTypes[i].Doc {
				result += fmt.Sprintf("%s\n", t.scannedTypes[i].Doc[j])
			}
		}

		if t.scannedTypes[i].InternalType == syntaxtree.ConstType {
			result += fmt.Sprintf("export const %s = %s;\n\n", t.scannedTypes[i].Name, t.scannedTypes[i].Value)
		}
	}

	for i := range t.scannedStructs {
		var extendsClasses []string
		for j := range t.scannedStructs[i].Fields {
			if isEmbeddedStructForInheritance(t.scannedStructs[i].Fields[j]) {
				extendsClasses = append(extendsClasses, t.scannedStructs[i].Fields[j].Name)
			}
		}

		result += fmt.Sprintf("\nexport interface %s", t.scannedStructs[i].Name)
		if len(extendsClasses) > 0 {
			result += fmt.Sprintf(" extends %s", strings.Join(extendsClasses, ", "))
		}

		result += fmt.Sprint(" {\n")

		for j := range t.scannedStructs[i].Fields {
			if isEmbeddedStructForInheritance(t.scannedStructs[i].Fields[j]) {
				continue
			}

			tag := CleanTagName(t.scannedStructs[i].Fields[j].Tag)
			if tag == "" || t.preserve {
				tag = t.scannedStructs[i].Fields[j].Name
			}

			result += fmt.Sprintf("\t%s: %s;\n", tag, getTypescriptCompatibleType(t.scannedStructs[i].Fields[j].Kind))
		}

		result += "}\n"
	}

	//fmt.Printf("%s", result)

	return result
}

func isEmbeddedStructForInheritance(field syntaxtree.ScannedStructField) bool {
	return field.Kind == "struct" && field.Tag == ""
}

func getTypescriptCompatibleType(goType string) string {
	result, ok := goTypeScriptTypeMappings[goType]
	if !ok {
		return goType
	}

	return result
}

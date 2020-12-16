package translators

import (
	"fmt"

	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
)

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
	fmt.Println("-----------------\nPerforming TypeScript translation!\n-----------------")

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
		result += fmt.Sprintf("\nexport interface %s {\n", t.scannedStructs[i].Name)
		for j := range t.scannedStructs[i].Fields {
			tag := CleanTagName(t.scannedStructs[i].Fields[j].Tag)
			if tag == "" || t.preserve {
				tag = t.scannedStructs[i].Fields[j].Name
			}

			result += fmt.Sprintf("\t%s: %s;\n", tag, t.scannedStructs[i].Fields[j].Kind)
		}

		result += "}\n"
	}

	fmt.Printf("%s", result)
	return result
}

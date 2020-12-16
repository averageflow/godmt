package translators

import (
	"fmt"

	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
)

type TypeScriptTranslator struct {
	scannedTypes   []syntaxtree.ScannedType
	scannedStructs []syntaxtree.ScannedStruct
}

func (t *TypeScriptTranslator) Setup(d []syntaxtree.ScannedType, s []syntaxtree.ScannedStruct) {
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
			result += fmt.Sprintf("export const %s = %s;\n", t.scannedTypes[i].Name, t.scannedTypes[i].Value)
		}
	}

	for i := range t.scannedStructs {
		result += fmt.Sprintf("\nexport interface %s  {\n", t.scannedStructs[i].Name)
		for j := range t.scannedStructs[i].Fields {
			result += fmt.Sprintf("\t%s: %s;\n", CleanTagName(t.scannedStructs[i].Fields[j].Tag), t.scannedStructs[i].Fields[j].Kind)
		}

		result += "}\n"
	}

	fmt.Printf("%s", result)
	return result
}

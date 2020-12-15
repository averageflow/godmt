package translators

import (
	"fmt"
	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
)

type TypeScriptTranslator struct {
	raw []syntaxtree.RawScannedType
}

func (t *TypeScriptTranslator) Setup(d []syntaxtree.RawScannedType) {
	t.raw = d
}

func (t *TypeScriptTranslator) Translate() string {
	fmt.Println("-----------------\nPerforming TypeScript translation!\n-----------------")
	//stringRepresentation := fmt.Sprintf("%+v", t.raw)
	//fmt.Printf("%+v\n", stringRepresentation)

	var result string

	for i := range t.raw{
		if t.raw[i].InternalType == syntaxtree.ConstType {
			result += fmt.Sprintf("export const %s = %s;\n", t.raw[i].Name, t.raw[i].Value)
		}
	}


	fmt.Printf("%s", result)
	return result
}

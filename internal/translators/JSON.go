package translators

import (
	"encoding/json"
	"fmt"
	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
)

type finalResult struct {
	Enums []syntaxtree.ScannedType   `json:"enums"`
	Types []syntaxtree.ScannedStruct `json:"types"`
}

type JSONTranslator struct {
	scannedTypes   []syntaxtree.ScannedType
	scannedStructs []syntaxtree.ScannedStruct
}

func (t *JSONTranslator) Setup(d []syntaxtree.ScannedType, s []syntaxtree.ScannedStruct) {
	t.scannedTypes = d
	t.scannedStructs = s
}

func (t *JSONTranslator) Translate() string {
	fmt.Println("-----------------\nPerforming JSON translation!\n-----------------")


	result, err := json.Marshal(finalResult{
		Enums: t.scannedTypes,
		Types: t.scannedStructs,
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s", result)
	return string(result)
}

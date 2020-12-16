package translators

import (
	"encoding/json"
	"fmt"

	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
)

type jsonFinalResult struct {
	Enums map[string]syntaxtree.ScannedType   `json:"enums"`
	Types map[string]syntaxtree.ScannedStruct `json:"types"`
}

type JSONTranslator struct {
	preserve       bool
	scannedTypes   map[string]syntaxtree.ScannedType
	scannedStructs map[string]syntaxtree.ScannedStruct
}

func (t *JSONTranslator) Setup(preserve bool, d map[string]syntaxtree.ScannedType, s map[string]syntaxtree.ScannedStruct) {
	t.preserve = preserve
	t.scannedTypes = d
	t.scannedStructs = s
}

func (t *JSONTranslator) Translate() string {
	fmt.Println(`
----------------------------------
Performing JSON translation!
----------------------------------
	`)

	payload := jsonFinalResult{
		Enums: t.scannedTypes,
		Types: t.scannedStructs,
	}

	result, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Printf("%s", result)

	return string(result)
}

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
	Translator
}

func (t *JSONTranslator) Setup(preserve bool, d map[string]syntaxtree.ScannedType, s map[string]syntaxtree.ScannedStruct) {
	t.Preserve = preserve
	t.ScannedTypes = d
	t.ScannedStructs = s
}

func (t *JSONTranslator) Translate() string {
	fmt.Println(`
----------------------------------
Performing JSON translation!
----------------------------------
	`)

	payload := jsonFinalResult{
		Enums: t.ScannedTypes,
		Types: t.ScannedStructs,
	}

	result, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Printf("%s", result)

	return string(result)
}

package translators

import (
	"encoding/json"
	"fmt"

	"github.com/averageflow/goschemaconverter/pkg/syntaxtreeparser"
)

type jsonFinalResult struct {
	Enums map[string]syntaxtreeparser.ScannedType   `json:"enums"`
	Types map[string]syntaxtreeparser.ScannedStruct `json:"types"`
}

type JSONTranslator struct {
	Translator
}

func (t *JSONTranslator) Setup(preserve bool, d map[string]syntaxtreeparser.ScannedType, s map[string]syntaxtreeparser.ScannedStruct) {
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

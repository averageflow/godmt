package translators

import (
	"encoding/json"
	"fmt"

	"github.com/averageflow/godmt/pkg/syntaxtreeparser"
)

type jsonFinalResult struct {
	Enums map[string]syntaxtreeparser.ScannedType   `json:"enums"`
	Types map[string]syntaxtreeparser.ScannedStruct `json:"types"`
}

type JSONTranslator struct {
	Translator
}

func (t *JSONTranslator) Translate() string {
	fmt.Println(`
----------------------------------
Performing JSON translation!
----------------------------------
	`)

	payload := jsonFinalResult{
		Enums: t.Data.ScanResult,
		Types: t.Data.StructScanResult,
	}

	result, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}

	return string(result)
}

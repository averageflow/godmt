package translators

import (
	"encoding/json"
	"fmt"

	"github.com/averageflow/godmt/pkg/godmt"
)

type jsonFinalResult struct {
	Enums map[string]godmt.ScannedType   `json:"enums"`
	Types map[string]godmt.ScannedStruct `json:"types"`
}

type JSONTranslator struct {
	Translator
}

func (t *JSONTranslator) Translate() string {
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

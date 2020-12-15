package translators

import (
	"encoding/json"
	"fmt"
	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
)

type JSONTranslator struct {
	raw []syntaxtree.RawScannedType
}

func (t *JSONTranslator) Setup(d []syntaxtree.RawScannedType) {
	t.raw = d
}

func (t *JSONTranslator) Translate() string {
	fmt.Println("-----------------\nPerforming JSON translation!\n-----------------")

	result, err := json.Marshal(t.raw)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("%s", result)
	return string(result)
}

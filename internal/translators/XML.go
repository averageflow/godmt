package translators

import (
	"encoding/xml"
	"fmt"

	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
)

type finalResult struct {
	Enums []syntaxtree.ScannedType   `xml:"enums"`
	Types []syntaxtree.ScannedStruct `xml:"types"`
}

type XMLTranslator struct {
	preserve       bool
	scannedTypes   []syntaxtree.ScannedType
	scannedStructs []syntaxtree.ScannedStruct
}

func (t *XMLTranslator) Setup(preserve bool, d []syntaxtree.ScannedType, s []syntaxtree.ScannedStruct) {
	t.preserve = preserve
	t.scannedTypes = d
	t.scannedStructs = s
}

func (t *XMLTranslator) Translate() string {
	fmt.Println(`
----------------------------------
Performing XML translation!
----------------------------------
	`)

	var typesWithoutMaps []syntaxtree.ScannedType

	for i := range t.scannedTypes {
		if t.scannedTypes[i].InternalType == syntaxtree.MapType {
			continue
		}
		typesWithoutMaps = append(typesWithoutMaps, t.scannedTypes[i])
	}

	payload := finalResult{
		Enums: typesWithoutMaps,
		Types: t.scannedStructs,
	}

	result, err := xml.MarshalIndent(payload, "", "\t")
	if err != nil {
		fmt.Println(err.Error())
	}

	//fmt.Printf("%s", result)

	return string(result)
}

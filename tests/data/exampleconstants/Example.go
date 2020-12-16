package exampleconstants

import "github.com/averageflow/goschemaconverter/tests/data/examplestructs"

const (
	NumberOne = 1
)

const (
	// Comment Test ExampleString
	ExampleString = "example"
)

const (
	//TrueConstant is true
	TrueConstant = true
)

type JustAType struct {
	Test string `json:"string"`
}

type NaughtyImport struct {
	DangerousField examplestructs.TestOne
	TestField      JustAType
}

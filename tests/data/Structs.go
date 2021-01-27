package data

import "io"

// ScannedType represents a basic entity to be translated.
// More specifically const and var items.
type ScannedType struct {
	Name         string      `json:"name"`
	Kind         io.Closer   `json:"kind"`
	KindSlice    []io.Closer `json:"kind_slice"`
	Value        interface{} `json:"value"`
	Doc          []string    `json:"doc"`
	InternalType int         `uri:"internalType"`
	Test         string
}

type ExtendedScannedType struct {
	ScannedType
	AnExtraField bool `json:"anExtraField"`
}

// ScannedStruct represents the details of a scanned struct.
type ScannedStruct struct {
	Doc          []string `json:"doc" binding:"required" validation:"required"`
	Name         string   `json:"name" binding:"required" validation:"required"`
	Fields       []bool   `json:"fields_test" binding:"required" validation:"required"`
	InternalType int      `xml:"internal_type" binding:"required" validation:"required"`
}

type InheritSlice struct {
	Structs         []ScannedStruct  `json:"structs"`
	PointerStructs  *[]ScannedStruct `json:"pointerStructs"`
	SliceOfPointers []*string        `json:"sliceOfPointers"`
}

type EmbeddedType struct {
	ID   string `json:"id"`
	Data struct {
		Test string `json:"test"`
	} `json:"data"`
}

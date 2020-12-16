package syntaxtree

const (
	MapType    = 1
	VarType    = 2
	ConstType  = 3
	StructType = 4
)

type ScannedType struct {
	Name         string      `json:"name" xml:""`
	Kind         string      `json:"kind" xml:""`
	Value        interface{} `json:"value" xml:""`
	Doc          []string    `json:"doc" xml:""`
	InternalType int         `json:"internalType" xml:""`
}

type ScannedStructField struct {
	Name string   `json:"name" xml:""`
	Kind string   `json:"kind" xml:""`
	Tag  string   `json:"tag" xml:""`
	Doc  []string `json:"doc" xml:""`
	//SubStruct *ScannedStruct `json:"subStruct"`
}

type ScannedStruct struct {
	Doc          []string             `json:"doc" xml:""`
	Name         string               `json:"name" xml:""`
	Fields       []ScannedStructField `json:"fields" xml:""`
	InternalType int                  `json:"internalType" xml:""`
}

type visitor int

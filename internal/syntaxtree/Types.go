package syntaxtree

const (
	MapType    = 1
	VarType    = 2
	ConstType  = 3
	StructType = 4
)

type ScannedType struct {
	Name         string      `json:"name"`
	Kind         string      `json:"kind"`
	Value        interface{} `json:"value"`
	Doc          []string    `json:"doc"`
	InternalType int         `json:"internalType"`
}


type ScannedStructField struct {
	Name string   `json:"name"`
	Kind string   `json:"kind"`
	Tag  string   `json:"tag"`
	Doc  []string `json:"doc"`
	//SubStruct *ScannedStruct `json:"subStruct"`
}

type ScannedStruct struct {
	Doc          []string             `json:"doc"`
	Name         string               `json:"name"`
	Fields       []ScannedStructField `json:"fields"`
	InternalType int                  `json:"internalType"`
}

type visitor int

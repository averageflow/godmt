package syntaxtreeparser

const (
	MapType    = 1
	VarType    = 2
	ConstType  = 3
	StructType = 4
	SliceType  = 5
)

type ScannedType struct {
	Name         string      `json:"name"`
	Kind         string      `json:"kind"`
	Value        interface{} `json:"value"`
	Doc          []string    `json:"doc"`
	InternalType int         `json:"internalType"`
}

type ImportedEntityDetails struct {
	EntityName  string
	PackageName string
}

type ScannedStructField struct {
	Name          string                 `json:"name"`
	Kind          string                 `json:"kind"`
	Tag           string                 `json:"tag"`
	Doc           []string               `json:"doc"`
	ImportDetails *ImportedEntityDetails `json:"imported_entity"`
}

type ScannedStruct struct {
	Doc          []string             `json:"doc"`
	Name         string               `json:"name"`
	Fields       []ScannedStructField `json:"fields"`
	InternalType int                  `json:"internalType"`
}

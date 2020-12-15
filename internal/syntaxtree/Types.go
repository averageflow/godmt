package syntaxtree

const (
	MapType   = 1
	VarType   = 2
	ConstType = 3
)

type RawScannedType struct {
	Name         string      `json:"name"`
	Kind         string      `json:"kind"`
	Value        interface{} `json:"value"`
	Doc          []string    `json:"doc"`
	InternalType int         `json:"internalType"`
}

type visitor int

var ScanResult []RawScannedType

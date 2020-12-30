package godmt

// The below constants represent the InternalType for a scanned type
// this is used in the translators to determine more easily the type of object we are
// performing operations on.
const (
	MapType    = 1
	VarType    = 2
	ConstType  = 3
	StructType = 4
	SliceType  = 5
)

// ScannedType represents a basic entity to be translated.
// More specifically const and var items.
type ScannedType struct {
	Name         string      `json:"name"`
	Kind         string      `json:"kind"`
	Value        interface{} `json:"value"`
	Doc          []string    `json:"doc"`
	InternalType int         `json:"internalType"`
}

// ImportedEntityDetails represents the details of an imported item.
// With that information an attempt of an import can be re-created in the translators.
type ImportedEntityDetails struct {
	EntityName  string
	PackageName string
}

// ScannedStructField represents the details of a field inside a scanned struct.
type ScannedStructField struct {
	Name          string                 `json:"name"`
	Kind          string                 `json:"kind"`
	Tag           string                 `json:"tag"`
	Doc           []string               `json:"doc"`
	ImportDetails *ImportedEntityDetails `json:"importedEntity"`
	SubFields     []ScannedStructField   `json:"subFields"`
	InternalType  int                    `json:"internalType"`
}

// ScannedStruct represents the details of a scanned struct.
type ScannedStruct struct {
	Doc          []string             `json:"doc"`
	Name         string               `json:"name"`
	Fields       []ScannedStructField `json:"fields"`
	InternalType int                  `json:"internalType"`
}

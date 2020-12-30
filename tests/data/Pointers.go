package data

// ScanPointers represents the details of a scanned struct.
type ScanPointers struct {
	Doc          *[]string       `json:"doc" binding:"required" validation:"required"`
	Name         *string         `json:"name" binding:"required" validation:"required"`
	Fields       *bool           `json:"fields" binding:"required" validation:"required"`
	InternalType *int            `xml:"internalType" binding:"required" validation:"required"`
	Test         *map[string]int `xml:"test"`
}

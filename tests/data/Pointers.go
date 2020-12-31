package data

// ScanPointers represents the details of a scanned struct.
type ScanPointers struct {
	// Doc deserves a comment
	Doc *[]string `json:"doc" binding:"required" validation:"required"`
	// Name also deserves a comment
	Name         *string                   `json:"name" binding:"required" validation:"required"`
	Fields       *bool                     `json:"fields" binding:"required" validation:"required"`
	InternalType *int                      `xml:"internalType" binding:"required" validation:"required"`
	Test         *map[string]int           `xml:"test"`
	MapOfSlices  *map[string][]int         `json:"tricky"`
	Insane       *map[int][]map[string]int `json:"insane"`
	MoreTest     *[]map[string]int         `xml:"moreTest"`
}

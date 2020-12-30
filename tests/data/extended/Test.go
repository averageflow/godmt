package extended

// ComplicatedMaps is a class that deserves a doc comment
type ComplicatedMaps struct {
	// StrangeMapOne is a complicated type
	StrangeMapOne map[string]map[string]int `json:"strangeMapOne"`
	// StrangeSliceOfMaps is also complicated type
	StrangeSliceOfMaps []map[int]string `json:"strangeSliceOfMaps"`
	SimpleMap          map[string]int   `json:"simpleMap"`
	MapOfSlices        map[string][]int `json:"tricky"`
}

// ScannedStruct represents the details of a scanned struct.
type ScannedStruct struct {
	Doc []string `json:"doc" binding:"required" validation:"required"`
	// Name also deserves a comment
	//Name         string `json:"name" binding:"required" validation:"required"`
	//Fields       []bool `json:"fields" binding:"required" validation:"required"`
	//InternalType int    `xml:"internalType" binding:"required" validation:"required"`
}

package data

// ComplicatedMaps is a class that deserves a doc comment.
type ComplicatedMaps struct {
	// StrangeMapOne is a complicated type
	StrangeMapOne map[string]map[string]int `json:"strangeMapOne"`
	// StrangeSliceOfMaps is also complicated type
	StrangeSliceOfMaps []map[int]string         `json:"strangeSliceOfMaps"`
	SimpleMap          map[string]int           `json:"simpleMap"`
	MapOfSlices        map[string][]int         `json:"tricky"`
	Insane             map[int][]map[string]int `json:"insane"`
}

package examplevars

var (
	// NumberOne is a nice example
	// that spans multiple comment lines
	NumberOneVar = 1
)

var (
	// ExampleString is a nice example
	ExampleStringVar = "example"
	// This map also deserves a comment
	MapTestVar = map[string]string{
		"test":  "example",
		"test2": "example2",
		"test3": "example3",
	}
	MapTestIntVar = map[string]int{
		"test":  1,
		"test2": 2,
		"test3": 3,
	}
)

var (
	//TrueConstant is true
	TrueVar = true
)

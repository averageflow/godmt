package data

var StringStringMap = map[string]string{ //nolint:gochecknoglobals
	"test":  "123",
	"test2": "234",
}

var IntIntMap = map[int]int{ //nolint:gochecknoglobals
	0: 123,
	1: 234,
}

var IntStringMap = map[int]string{ //nolint:gochecknoglobals
	0: "test",
	1: "test1",
}

var StringIntMap = map[string]int{ //nolint:gochecknoglobals
	"test":  123,
	"test1": 234,
}

var StringAnyMap = map[string]interface{}{ //nolint:gochecknoglobals
	"test a string": "test",
	"test a int":    123,
	"test a bool":   false,
}

var StringBoolMap = map[string]bool{ //nolint:gochecknoglobals
	"godmt": true,
	"typex": false,
}

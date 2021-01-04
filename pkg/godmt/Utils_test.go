package godmt

import "testing"

func TestCleanTagName(t *testing.T) {
	t.Parallel()

	testTable := map[string]string{
		"`json:\"testField\"`":                      "testField",
		"`json:\"testField\" binding:\"required\"`": "testField",
		"`xml:\"testField\"`":                       "testField",
		"`form:\"testField\"`":                      "testField",
		"`uri:\"testField\"`":                       "testField",
		"`db:\"testField\"`":                        "testField",
		"`header:\"testField\"`":                    "testField",
		"`mapstructure:\"testField\"`":              "testField",
		"":                                          "",
		"whatever:\"testField\"":                    "",
	}

	for i := range testTable {
		cleanTag := CleanTagName(i)
		if cleanTag != testTable[i] {
			t.Errorf("Expected %s, got %s", testTable[i], cleanTag)
		}
	}
}

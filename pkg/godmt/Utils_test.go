package godmt

import (
	"go/ast"
	"testing"
)

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

func TestSliceValuesToPrettyList(t *testing.T) {
	t.Parallel()

	expected := "\ttest1,\n\ttest2,\n\ttest3"
	sut := SliceValuesToPrettyList([]string{"test1", "test2", "test3"})

	if sut != expected {
		t.Errorf("Expected %s, got %s", expected, sut)
	}
}

func TestExtractComments(t *testing.T) {
	t.Parallel()

	expected := []string{"// test1", "// test2"}

	commentGroup := &ast.CommentGroup{
		List: []*ast.Comment{
			{
				Slash: 0,
				Text:  "// test1",
			},
			{
				Slash: 0,
				Text:  "// test2",
			},
		},
	}

	sut := ExtractComments(commentGroup)

	if sut[0] != expected[0] {
		t.Errorf("Expected %s, got %s", expected[0], sut[0])
	}

	if sut[1] != expected[1] {
		t.Errorf("Expected %s, got %s", expected[1], sut[1])
	}
}

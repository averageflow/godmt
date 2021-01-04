package utils

import "testing"

func TestGetFileDestination(t *testing.T) {
	t.Parallel()

	testTable := [][]string{
		{
			"./result/Example.go", "./result", "../../tests/data/Example.go", "../../tests/data",
		},
		{
			"./result/Example.go", "./result", "../../tests/data/Example.go", "../../tests/data/Example.go",
		},
	}

	for i := range testTable {
		sut := GetFileDestination(testTable[i][1], testTable[i][2], testTable[i][3])
		if sut != testTable[i][0] {
			t.Errorf("Expected %s, got %s", testTable[i][0], sut)
		}
	}
}

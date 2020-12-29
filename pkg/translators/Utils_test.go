package translators

import (
	"testing"

	"github.com/averageflow/godmt/pkg/godmt"
)

func TestIsEmbeddedStructForInheritance(t *testing.T) {
	sut := godmt.ScannedStructField{
		Name:          "",
		Kind:          "",
		Tag:           "",
		Doc:           nil,
		ImportDetails: nil,
		SubFields:     nil,
	}

	if IsEmbeddedStructForInheritance(&sut) {
		t.Errorf("Expected sut to not be an embedded struct for inheritance")
	}

	sut.Kind = StructTypeKeyWord
	sut.Tag = ""

	if !IsEmbeddedStructForInheritance(&sut) {
		t.Errorf("Expected sut to be an embedded struct for inheritance")
	}
}

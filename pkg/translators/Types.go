package translators

import (
	"github.com/averageflow/godmt/internal/syntaxtree"
	"github.com/averageflow/godmt/pkg/godmt"
)

const (
	TypeScriptTranslationMode = "typescript"
	TSTranslationMode         = "ts"
	SwiftTranslationMode      = "swift"
	JSONTranslationMode       = "json"
)

type Translator struct {
	Preserve bool
	Data     syntaxtree.FileResult
}

type TypeTranslator interface {
	Translate() string
	Setup(d []godmt.ScannedType)
}

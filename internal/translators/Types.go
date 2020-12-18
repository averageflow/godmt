package translators

import (
	"github.com/averageflow/goschemaconverter/pkg/syntaxtreeparser"
)

const (
	TypeScriptTranslationMode = "typescript"
	TSTranslationMode         = "ts"
	SwiftTranslationMode      = "swift"
	JSONTranslationMode       = "json"
)

type Translator struct {
	Preserve       bool
	OrderedTypes   []string
	ScannedTypes   map[string]syntaxtreeparser.ScannedType
	OrderedStructs []string
	ScannedStructs map[string]syntaxtreeparser.ScannedStruct
}

type TypeTranslator interface {
	Translate() string
	Setup(d []syntaxtreeparser.ScannedType)
}

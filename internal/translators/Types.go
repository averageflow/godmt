package translators

import "github.com/averageflow/goschemaconverter/internal/syntaxtree"

const (
	TypeScriptTranslationMode = "typescript"
	TSTranslationMode         = "ts"
	SwiftTranslationMode      = "swift"
)

type Translator struct {
	Preserve       bool
	OrderedTypes   []string
	ScannedTypes   map[string]syntaxtree.ScannedType
	OrderedStructs []string
	ScannedStructs map[string]syntaxtree.ScannedStruct
}

type TypeTranslator interface {
	Translate() string
	Setup(d []syntaxtree.ScannedType)
}

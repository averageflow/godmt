package translators

import "github.com/averageflow/goschemaconverter/internal/syntaxtree"

const (
	TypeScriptTranslationMode = "typescript"
	TSTranslationMode         = "ts"
	SwiftTranslationMode      = "swift"
)

type TypeTranslator interface {
	Translate() string
	Setup(d []syntaxtree.ScannedType)
}

package translators

import "github.com/averageflow/goschemaconverter/internal/syntaxtree"

const (
	TypeScriptTranslationMode = "typescript"
	TSTranslationMode         = "ts"
)

type TypeTranslator interface {
	Translate() string
	Setup(d []syntaxtree.ScannedType)
}

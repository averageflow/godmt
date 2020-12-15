package translators

import "github.com/averageflow/goschemaconverter/internal/syntaxtree"

type TypeTranslator interface {
	Translate() string
	Setup(d []syntaxtree.RawScannedType)
}

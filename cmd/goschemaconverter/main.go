package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
	"github.com/averageflow/goschemaconverter/internal/translators"
)

func main() {
	scanPath := flag.String("dir", ".", "directory to scan")
	translateMode := flag.String("translation", "json", "translation mode")
	preserveNames := flag.Bool("preserve", false, "should preserve the original struct field names")

	flag.Parse()

	err := filepath.Walk(*scanPath, syntaxtree.ScanDir)
	if err != nil {
		log.Println(err)
	}

	switch *translateMode {
	case translators.JSONTranslationMode:
		j := translators.JSONTranslator{}
		j.Setup(*preserveNames, syntaxtree.ScanResult, syntaxtree.StructScanResult)
		j.Translate()
	case translators.TypeScriptTranslationMode:
	case translators.TSTranslationMode:
		ts := translators.TypeScriptTranslator{}
		ts.Setup(*preserveNames, syntaxtree.ScanResult, syntaxtree.StructScanResult)
		ts.Translate()
	}
}

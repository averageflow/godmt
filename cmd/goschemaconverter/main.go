package main

import (
	"flag"
	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
	"github.com/averageflow/goschemaconverter/internal/translators"
	"log"
	"path/filepath"
)


func main() {
	scanPath := flag.String("dir", ".", "directory to scan")
	translateMode := flag.String("translation", "json", "translation mode")

	flag.Parse()

	err := filepath.Walk(*scanPath, syntaxtree.ScanDir)
	if err != nil {
		log.Println(err)
	}

	switch *translateMode {
	case translators.JSONTranslationMode:
		j := translators.JSONTranslator{}
		j.Setup(syntaxtree.ScanResult, syntaxtree.StructScanResult)
		j.Translate()
		break
	case translators.TypeScriptTranslationMode:
	case translators.TSTranslationMode:
		ts := translators.TypeScriptTranslator{}
		ts.Setup(syntaxtree.ScanResult, syntaxtree.StructScanResult)
		ts.Translate()
		break
	}
}
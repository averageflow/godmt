package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
	"github.com/averageflow/goschemaconverter/internal/translators"
)

func main() {
	scanPath := flag.String("dir", ".", "directory to scan")
	translateMode := flag.String("translation", "json", "translation mode")
	preserveNames := flag.Bool("preserve", false, "should preserve the original struct field names")

	flag.Parse()

	err := filepath.Walk(*scanPath, syntaxtree.GetFileCount)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Processing %d files!\n", syntaxtree.TotalFileCount)

	err = filepath.Walk(*scanPath, syntaxtree.ScanDir)
	if err != nil {
		log.Println(err)
	}

	var resultingOutput string
	var filename string

	switch *translateMode {

	case translators.TypeScriptTranslationMode:
	case translators.TSTranslationMode:
		filename = "result.ts"
		ts := translators.TypeScriptTranslator{}
		ts.Setup(*preserveNames, syntaxtree.ScanResult, syntaxtree.StructScanResult)
		resultingOutput = ts.Translate()
	case translators.XMLTranslationMode:
		filename = "result.xml"
		x := translators.XMLTranslator{}
		x.Setup(*preserveNames, syntaxtree.ScanResult, syntaxtree.StructScanResult)
		resultingOutput = x.Translate()
	default:
		filename = "result.json"
		j := translators.JSONTranslator{}
		j.Setup(*preserveNames, syntaxtree.ScanResult, syntaxtree.StructScanResult)
		resultingOutput = j.Translate()
	}

	f, err := os.Create(filename)

	defer f.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	l, err := f.WriteString(resultingOutput)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}

	fmt.Printf("%d bytes written successfully", l)
}

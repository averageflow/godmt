package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/averageflow/godmt/pkg/translators"
)

type OperationMode struct {
	WantedPath    string
	TranslateMode string
	PreserveNames bool
	Tree          bool
}

func ParseApplicationConfig() OperationMode {
	scanPath := flag.String("dir", "", "directory to scan")
	translateMode := flag.String("translation", translators.JSONTranslationMode, "translation mode")
	preserveNames := flag.Bool("preserve", false, "should preserve the original struct field names")
	tree := flag.Bool("tree", false, "should show the abstract syntax tree")

	flag.Parse()

	wantedPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err.Error())
	}

	if *scanPath != "" {
		wantedPath = *scanPath
	} else {
		fmt.Printf("No directory specified! Defaulting to current working directory:\n%s\n", wantedPath)
	}

	wantedPath = strings.TrimSuffix(wantedPath, string(os.PathSeparator))

	return OperationMode{
		WantedPath:    wantedPath,
		TranslateMode: *translateMode,
		PreserveNames: *preserveNames,
		Tree:          *tree,
	}
}

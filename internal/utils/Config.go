package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type OperationMode struct {
	WantedPath    string
	TranslateMode string
	PreserveNames bool
	Tree          bool
	Destination   string
}

func ParseApplicationConfig() OperationMode {
	scanPath := flag.String("dir", "", "directory to scan")
	translateMode := flag.String("translation", "json", "translation mode")
	preserveNames := flag.Bool("preserve", false, "should preserve the original struct field names")
	tree := flag.Bool("tree", false, "should show the abstract syntax tree")
	destination := flag.String("o", fmt.Sprintf(".%sresult", string(os.PathSeparator)), "destination of output")

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
		WantedPath:    filepath.Clean(wantedPath),
		TranslateMode: *translateMode,
		PreserveNames: *preserveNames,
		Tree:          *tree,
		Destination:   filepath.Clean(*destination),
	}
}

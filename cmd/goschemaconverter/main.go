package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/averageflow/goschemaconverter/internal/syntaxtree"
	"github.com/averageflow/goschemaconverter/internal/translators"
)

func main() {
	scanPath := flag.String("dir", ".", "directory to scan")
	translateMode := flag.String("translation", translators.JSONTranslationMode, "translation mode")
	preserveNames := flag.Bool("preserve", false, "should preserve the original struct field names")
	tree := flag.Bool("tree", false, "should show the abstract syntax tree")
	flag.Parse()

	syntaxtree.Result = make(map[string]syntaxtree.FileResult)
	syntaxtree.ShouldPrintAbstractSyntaxTree = *tree

	err := filepath.Walk(*scanPath, syntaxtree.GetFileCount)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Processing %d files!\n", syntaxtree.TotalFileCount)

	err = filepath.Walk(*scanPath, syntaxtree.ScanDir)
	if err != nil {
		log.Println(err)
	}

	if syntaxtree.ShouldPrintAbstractSyntaxTree {
		fmt.Printf("0 bytes written")
		os.Exit(0)
	}

	for i := range syntaxtree.Result {
		item := syntaxtree.Result[i]
		constantsOrderSlice := make([]string, len(item.ScanResult))
		structsOrderSlice := make([]string, len(item.StructScanResult))

		j := 0
		for k := range item.ScanResult {
			constantsOrderSlice[j] = k
			j++
		}

		j = 0
		for k := range item.StructScanResult {
			structsOrderSlice[j] = k
			j++
		}

		sort.Strings(constantsOrderSlice)
		sort.Strings(structsOrderSlice)

		item.ConstantSort = constantsOrderSlice
		item.StructSort = structsOrderSlice
		syntaxtree.Result[i] = item
	}

	os.Mkdir("./result", os.FileMode(0777))
	var resultingOutput string

	for i := range syntaxtree.Result {
		baseTranslator := translators.Translator{
			Preserve: *preserveNames,
			Data:     syntaxtree.Result[i],
		}

		cleanPath := strings.TrimSuffix(*scanPath, "/")

		filename := fmt.Sprintf("./result%s", strings.ReplaceAll(i, cleanPath, ""))

		packageName := strings.Split(strings.ReplaceAll(i, cleanPath, ""), "/")

		if len(packageName) == 3 {
			os.Mkdir(fmt.Sprintf("./result/%s", packageName[1]), os.FileMode(0777))
		}

		switch *translateMode {

		case translators.TypeScriptTranslationMode:
		case translators.TSTranslationMode:
			filename = strings.ReplaceAll(filename, ".go", ".ts")
			ts := translators.TypeScriptTranslator{
				Translator: baseTranslator,
			}
			resultingOutput = ts.Translate()
			break
		case translators.SwiftTranslationMode:
			filename = strings.ReplaceAll(filename, ".go", ".swift")
			s := translators.SwiftTranslator{
				Translator: baseTranslator,
			}
			resultingOutput = s.Translate()
			break
		default:
			filename = strings.ReplaceAll(filename, ".go", ".json")
			j := translators.JSONTranslator{
				Translator: baseTranslator,
			}
			resultingOutput = j.Translate()
			break
		}

		f, err := os.Create(filename)

		defer f.Close()

		if err != nil {
			fmt.Printf("FILE OPERATION ERROR ON %s: %s\n", filename, err.Error())
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
}

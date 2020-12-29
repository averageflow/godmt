package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/averageflow/godmt/internal/syntaxtree"
	"github.com/averageflow/godmt/internal/translators"
	"github.com/averageflow/godmt/internal/utils"
	"github.com/cheggaaa/pb/v3"
)

func main() {
	syntaxtree.Result = make(map[string]syntaxtree.FileResult)

	config := utils.ParseApplicationConfig()

	err := filepath.Walk(config.WantedPath, syntaxtree.GetFileCount)
	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Processing %d files!\n", syntaxtree.TotalFileCount)

	err = filepath.Walk(config.WantedPath, syntaxtree.ScanDir)
	if err != nil {
		log.Println(err)
	}

	if syntaxtree.ShouldPrintAbstractSyntaxTree {
		// Exit by displaying the AST tree
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

	utils.CreateResultFolder()

	var resultingOutput string

	bar := pb.StartNew(syntaxtree.TotalFileCount)

	for i := range syntaxtree.Result {
		bar.Increment()

		baseTranslator := translators.Translator{
			Preserve: config.PreserveNames,
			Data:     syntaxtree.Result[i],
		}

		filename := fmt.Sprintf("./result%s", strings.ReplaceAll(i, config.WantedPath, ""))

		switch config.TranslateMode {

		case translators.TypeScriptTranslationMode:
		case translators.TSTranslationMode:
			filename = strings.ReplaceAll(filename, ".go", ".ts")
			ts := translators.TypeScriptTranslator{
				Translator: baseTranslator,
			}
			resultingOutput = ts.Translate()

		case translators.SwiftTranslationMode:
			filename = strings.ReplaceAll(filename, ".go", ".swift")
			s := translators.SwiftTranslator{
				Translator: baseTranslator,
			}
			resultingOutput = s.Translate()

		default:
			filename = strings.ReplaceAll(filename, ".go", ".json")
			j := translators.JSONTranslator{
				Translator: baseTranslator,
			}
			resultingOutput = j.Translate()

		}

		packageName := strings.Split(strings.ReplaceAll(i, config.WantedPath, ""), "/")
		utils.WriteResultToFile(resultingOutput, filename, packageName)
	}

	bar.Finish()
}

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/averageflow/godmt/internal/syntaxtree"
	"github.com/averageflow/godmt/internal/utils"
	"github.com/averageflow/godmt/pkg/translators"
	"github.com/cheggaaa/pb/v3"
)

func main() {
	config := utils.ParseApplicationConfig()
	syntaxtree.ResetGlobals(config)

	err := filepath.Walk(config.WantedPath, syntaxtree.GetFileCount)
	if err != nil {
		log.Println(err)
	}

	err = filepath.Walk(config.WantedPath, syntaxtree.ScanDir)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Parsed files successfully!")

	if config.Tree {
		// Exit by displaying the AST tree
		os.Exit(0)
	}

	for i := range syntaxtree.Result {
		syntaxtree.Result[i] = syntaxtree.GetOrderedFileItems(syntaxtree.Result[i])
	}

	fmt.Println("Sorted parsed items successfully!")

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
			filename = strings.ReplaceAll(filename, ".go", ".ts")
			ts := translators.TypeScriptTranslator{
				Translator: baseTranslator,
			}
			resultingOutput = ts.Translate()
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
		case translators.PHPTranslationMode:
			filename = strings.ReplaceAll(filename, ".go", ".php")
			s := translators.PHPTranslator{
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

		packageName := strings.Split(strings.ReplaceAll(i, config.WantedPath, ""), string(os.PathSeparator))
		utils.WriteResultToFile(resultingOutput, filename, packageName)
	}

	bar.Finish()
	fmt.Println("Translation was successful! All operations completed!")
}

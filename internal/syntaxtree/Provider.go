package syntaxtree

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"sort"
	"strings"

	"github.com/averageflow/godmt/pkg/godmt"
)

type visitor int

type FileResult struct {
	ConstantSort     []string
	ScanResult       map[string]godmt.ScannedType
	StructScanResult map[string]godmt.ScannedStruct
	StructSort       []string
}

var Result map[string]FileResult
var CurrentFile string
var TotalFileCount int
var ShouldPrintAbstractSyntaxTree bool

func WalkSyntaxTree(f *ast.File) {
	var v visitor

	ast.Walk(v, f)
}

// Visit represents the actions to be performed on every node of the tree
// n represents the node, whose type can be obtained with fmt.Sprintf and %T.
func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}

	if ShouldPrintAbstractSyntaxTree {
		fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
		fmt.Printf("%d", v)

		return v + 1
	}

	switch d := n.(type) {
	case *ast.Ident:
		if d.Obj == nil {
			return v + 1
		}

		if d.Obj.Kind == ast.Typ {
			result := parseStruct(d)

			for i := range result {
				_, ok := Result[CurrentFile].StructScanResult[result[i].Name]
				if !ok {
					Result[CurrentFile].StructScanResult[result[i].Name] = result[i]
				}
			}
		} else if d.Obj.Kind == ast.Con || d.Obj.Kind == ast.Var {
			result := parseConstantsAndVariables(d)

			for i := range result {
				_, ok := Result[CurrentFile].ScanResult[result[i].Name]
				if !ok {
					Result[CurrentFile].ScanResult[result[i].Name] = result[i]
				}
			}
		}

	default:
		break
	}

	return v + 1
}

func GetFileCount(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	TotalFileCount++

	return nil
}

func ScanDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	if !strings.Contains(info.Name(), ".go") {
		return nil
	}

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	CurrentFile = path
	currentResult := Result[path]

	currentResult.StructScanResult = make(map[string]godmt.ScannedStruct)
	currentResult.ScanResult = make(map[string]godmt.ScannedType)

	Result[path] = currentResult

	WalkSyntaxTree(f)

	return nil
}

func GetOrderedFileItems(item FileResult) FileResult {
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

	return item
}

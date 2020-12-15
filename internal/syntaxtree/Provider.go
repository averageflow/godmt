package syntaxtree

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
)

func WalkSyntaxTree(f *ast.File) {
	var v visitor

	ast.Walk(v, f)
}

var ScanResult []ScannedType
var StructScanResult []ScannedStruct

// Visit represents the actions to be performed on every node of the tree
// n represents the node, whose type can be obtained with fmt.Sprintf and %T
func (v visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	switch d := n.(type) {

	case *ast.Ident:
		if d.Obj == nil {
			return v + 1
		}

		if d.Obj.Kind == ast.Typ {
			fmt.Printf("type: %T, value: %+v\n", d.Obj, d.Obj)
			result := parseStruct(d)

			for i := range result {
				StructScanResult = append(StructScanResult, result[i])
			}
		}

		if d.Obj.Kind == ast.Con || d.Obj.Kind == ast.Var {
			result := parseConstantsAndVariables(d)

			for i := range result {
				ScanResult = append(ScanResult, result[i])
			}
		}

		break
	}
	return v + 1
}

func ScanDir(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if info.IsDir() {
		return nil
	}

	fmt.Printf("Scanning file: %s\n", path)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	WalkSyntaxTree(f)
	return nil
}

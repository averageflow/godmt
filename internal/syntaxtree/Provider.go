package syntaxtree

import (
	"go/ast"
)

func WalkSyntaxTree(f *ast.File) {
	var v visitor

	ast.Walk(v, f)
}


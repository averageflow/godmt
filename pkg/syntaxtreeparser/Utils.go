package syntaxtreeparser

import (
	"fmt"
	"go/ast"
	"strings"
)

func ExtractComments(rawCommentGroup *ast.CommentGroup) []string {
	var result []string
	if rawCommentGroup == nil {
		return result
	}

	commentList := rawCommentGroup.List
	for i := range commentList {
		result = append(result, commentList[i].Text)
	}

	return result
}

func GetMapValueType(item ast.Expr) string {
	switch item.(type) {
	case *ast.Ident:
		return item.(*ast.Ident).Name
	default:
		return "interface{}"
	}
}

func ExtractSliceValues(items []ast.Expr) []string {
	var result []string
	for i := range items {
		result = append(result, items[i].(*ast.BasicLit).Value)
	}
	return result
}

func SliceValuesToPrettyList(raw []string) string {
	var result []string

	for i := range raw {
		result = append(result, fmt.Sprintf("\t%s", raw[i]))
	}

	return strings.Join(result, ",\n")
}

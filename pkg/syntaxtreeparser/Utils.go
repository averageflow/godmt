package syntaxtreeparser

import (
	"fmt"
	"go/ast"
	"strings"
)

// ExtractComments will transform a *ast.CommentGroup into a []string
// which makes it more accessible and usable.
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

// GetMapValueType will return the type of a map's value fields.
func GetMapValueType(item ast.Expr) string {
	switch item.(type) {
	case *ast.Ident:
		return item.(*ast.Ident).Name
	default:
		return "interface{}"
	}
}

// ExtractSliceValues will return the values of a []ast.Expr in the form
// of a []string for ease of use.
func ExtractSliceValues(items []ast.Expr) []string {
	var result []string
	for i := range items {
		result = append(result, items[i].(*ast.BasicLit).Value)
	}
	return result
}

// SliceValuesToPrettyList will turn a normal []string into a line & comma
// separated string, for pretty display of a slice's values.
func SliceValuesToPrettyList(raw []string) string {
	var result []string

	for i := range raw {
		result = append(result, fmt.Sprintf("\t%s", raw[i]))
	}

	return strings.Join(result, ",\n")
}

func CleanTagName(rawTagName string) string {
	replacePatterns := []string{
		",string",
		"`json:\"",
		"`form:\"",
		"\" binding:\"",
		"`uri:\"",
		",omitempty",
		"\"`",
		`binding:"required"`,
	}

	result := rawTagName

	for i := range replacePatterns {
		result = strings.ReplaceAll(result, replacePatterns[i], "")
	}

	return strings.TrimSpace(result)
}

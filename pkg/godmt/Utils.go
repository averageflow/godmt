package godmt

import (
	"fmt"
	"go/ast"
	"strings"
)

// ExtractComments will transform a *ast.CommentGroup into a []string
// which makes it more accessible and usable.
func ExtractComments(rawCommentGroup *ast.CommentGroup) []string {
	if rawCommentGroup == nil {
		return nil
	}

	result := make([]string, len(rawCommentGroup.List))
	commentList := rawCommentGroup.List

	for i := range commentList {
		result[i] = commentList[i].Text
	}

	return result
}

// GetMapValueType will return the type of a map's value fields.
func GetMapValueType(item ast.Expr) string {
	switch value := item.(type) {
	case *ast.Ident:
		return value.Name
	default:
		return "interface{}"
	}
}

// ExtractSliceValues will return the values of a []ast.Expr in the form
// of a []string for ease of use.
func ExtractSliceValues(items []ast.Expr) []string {
	result := make([]string, len(items))
	for i := range items {
		switch item := items[i].(type) {
		case *ast.BasicLit:
			result[i] = item.Value
		case *ast.Ident:
			result[i] = item.Name
		default:
			break
		}
	}

	return result
}

// SliceValuesToPrettyList will turn a normal []string into a line & comma
// separated string, for pretty display of a slice's values.
func SliceValuesToPrettyList(raw []string) string {
	result := make([]string, len(raw))
	for i := range raw {
		result[i] = fmt.Sprintf("\t%s", raw[i])
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

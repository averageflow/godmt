package godmt

import (
	"fmt"
	"go/ast"
	"regexp"
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
	jsonRegex := regexp.MustCompile(`(?m)json:"[^"]+"`)
	xmlRegex := regexp.MustCompile(`(?m)xml:"[^"]+"`)

	var jsonTagName []string
	var xmlTagName []string

	// By default use JSON tag name
	for _, match := range jsonRegex.FindAllString(rawTagName, -1) {
		name := strings.TrimSpace(match)
		name = strings.ReplaceAll(name, `json:"`, ``)
		name = strings.ReplaceAll(name, `"`, ``)
		jsonTagName = append(jsonTagName, name)
	}

	if len(jsonTagName) > 0 {
		return strings.Join(jsonTagName, "")
	}

	// Alternatively use XML tag name
	for _, match := range xmlRegex.FindAllString(rawTagName, -1) {
		name := strings.TrimSpace(match)
		name = strings.ReplaceAll(name, `xml:"`, ``)
		name = strings.ReplaceAll(name, `"`, ``)
		xmlTagName = append(xmlTagName, name)
	}

	if len(xmlTagName) > 0 {
		return strings.Join(xmlTagName, "")
	}

	// Fallback to the raw tag
	return rawTagName
}

package translators

import (
	"strings"
)

func CleanTagName(rawTagName string) string {
	var result string

	result = strings.ReplaceAll(rawTagName, ",string", "")
	result = strings.ReplaceAll(result, "`json:\"", "")
	result = strings.ReplaceAll(result, ",omitempty", "")
	result = strings.ReplaceAll(result, "\"`", "")
	result = strings.ReplaceAll(result, `binding:"required"`, ``)
	return strings.TrimSpace(result)
}

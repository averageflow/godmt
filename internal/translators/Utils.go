package translators

import "strings"

func CleanTagName(rawTagName string) string {
	return strings.ReplaceAll(rawTagName, `json:"`, ``)
}

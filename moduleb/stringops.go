package moduleb

import (
	"strings"
	"unicode"
)

// ToCaps capitalises the given string
func ToCaps(s string) string {
	var sb strings.Builder

	for _, c := range s {
		sb.WriteRune(unicode.ToUpper(c))
	}

	return sb.String()
}

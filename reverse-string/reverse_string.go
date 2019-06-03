package reverse

import (
	"strings"
)

// String stuff
func String(in string) string {
	chars := []rune(in)
	var out strings.Builder
	for i := len(chars) - 1; i >= 0; i-- {
		out.WriteRune(chars[i])
	}
	return out.String()
}

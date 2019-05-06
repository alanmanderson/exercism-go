// generates an acronym from an input string
package acronym

import (
	"strings"
	"unicode"
)

// Takes a string and returns the acronym of all words
func Abbreviate(s string) string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && c != '\''
	}
	var acronym strings.Builder
	for _, element := range strings.FieldsFunc(s, f) {
		acronym.WriteString(strings.ToUpper(element[0:1]))
	}
	return acronym.String()
}

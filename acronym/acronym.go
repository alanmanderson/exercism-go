// Package acronym generates an acronym from an input string
package acronym

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// Abbreviate Takes a string and returns the acronym of all words
func Abbreviate(s string) string {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && c != '\''
	}
	var acronym strings.Builder
	for _, element := range strings.FieldsFunc(s, f) {
		r, _ := utf8.DecodeRuneInString(element)
		acronym.WriteString(strings.ToUpper(string(r)))
	}
	return acronym.String()
}

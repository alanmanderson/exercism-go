// Package acronym generates an acronym from an input string
package acronym

import (
	"strings"
	"unicode"
)

// Abbreviate Takes a string and returns the acronym of all words
func Abbreviate(s string) string {
	var acronym strings.Builder
	addLetter := true
	for _, r := range s {
		if !unicode.IsLetter(r) && r != '\'' {
			addLetter = true
			continue
		}
		if addLetter {
			acronym.WriteString(strings.ToUpper(string(r)))
		}
		addLetter = false
	}
	return acronym.String()
}

package wordcount

import (
	"strings"
	"unicode"
)

// Frequency stuff
type Frequency map[string]int

// WordCount stuff
func WordCount(in string) Frequency {
	out := make(Frequency, 0)
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c) && c != '\''
	}
	for _, word := range strings.FieldsFunc(in, f) {
		out[strings.ToLower(word)]++
	}
	return out
}

package anagram

import (
	"strings"
	"unicode"
)

// Detect stuff
func Detect(in string, anagrams []string) []string {
	out := make([]string, 0)
	for _, anagram := range anagrams {
		if areAnagrams(in, anagram) {
			out = append(out, anagram)
		}
	}
	return out
}

func areAnagrams(in1 string, in2 string) bool {
	if len(in1) != len(in2) || strings.EqualFold(in1, in2) {
		return false
	}
	dict := make(map[rune]int)
	for _, r := range in1 {
		dict[unicode.ToLower(r)]++
	}
	for _, r := range in2 {
		r = unicode.ToLower(r)
		dict[r]--
		if dict[r] < 0 {
			return false
		}
	}
	return true
}

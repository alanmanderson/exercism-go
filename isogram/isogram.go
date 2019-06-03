package isogram

import "unicode"

// IsIsogram tests if the input is an isogram
func IsIsogram(in string) bool {
	used := make(map[rune]bool)
	for _, char := range in {
		char = unicode.ToLower(char)
		if used[char] && unicode.IsLetter(char) {
			return false
		}
		used[char] = true
	}
	return true
}

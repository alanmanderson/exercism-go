package rotationalcipher

import (
	"strings"
	"unicode"
)

// RotationalCipher stuff
func RotationalCipher(text string, offset int) string {
	var sb strings.Builder
	for _, r := range text {
		if !unicode.IsLetter(r) {
			sb.WriteRune(r)
			continue
		}
		newRune := rune(int(r) + offset)
		if r <= 'Z' && newRune > 'Z' || (r <= 'z' && newRune > 'z') {
			newRune = rune(int(newRune) - 26)
		}

		sb.WriteRune(newRune)
	}
	return sb.String()
}

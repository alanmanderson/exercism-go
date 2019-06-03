package isbn

import (
	"strings"
	"unicode"
)

// IsValidISBN determines if the string is a valid isbn
func IsValidISBN(in string) bool {
	in = strings.Replace(in, "-", "", -1)
	if len(in) != 10 {
		return false
	}
	sum := 0
	for pos, r := range in {
		var digit int
		if pos == 9 && r == 'X' {
			digit = 10
		} else {
			if !unicode.IsNumber(r) {
				return false
			}
			digit = int(r - '0')
		}
		sum += (10 - pos) * digit
	}
	return sum%11 == 0
}

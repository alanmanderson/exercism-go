package pangram

import "unicode"

// IsPangram tests if a string is a pangram
func IsPangram(in string) bool {
	used := make(map[rune]bool)
	for _, r := range in {
		used[unicode.ToLower(r)] = true
	}
	return used['a'] && used['b'] && used['c'] && used['d'] && used['e'] && used['f'] && used['g'] && used['h'] && used['i'] && used['j'] && used['k'] && used['l'] && used['m'] && used['n'] && used['o'] && used['p'] && used['q'] && used['r'] && used['s'] && used['t'] && used['u'] && used['v'] && used['w'] && used['x'] && used['y'] && used['z']
}

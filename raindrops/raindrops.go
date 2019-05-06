package raindrops

import "strconv"

// Convert changes an integer into raindrop language
func Convert(input int) string {
	out := ""
	if input%3 == 0 {
		out += "Pling"
	}
	if input%5 == 0 {
		out += "Plang"
	}
	if input%7 == 0 {
		out += "Plong"
	}
	if out == "" {
		out = strconv.Itoa(input)
	}
	return out
}

package etl

import "strings"

// Transform handles the transform step of ETL
func Transform(in map[int][]string) (out map[string]int) {
	out = make(map[string]int)
	for i, arr := range in {
		for _, char := range arr {
			out[strings.ToLower(string(char))] = i
		}
	}
	return out
}

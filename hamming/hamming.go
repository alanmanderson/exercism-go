package hamming

import (
	"errors"
)

// Distance Calculate the hamming distance between two strings
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Lengths do not match")
	}

	difference := 0
	for pos := range a {
		if a[pos] != b[pos] {
			difference++
		}
	}
	return difference, nil
}

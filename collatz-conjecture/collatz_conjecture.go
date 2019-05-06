package collatzconjecture

import "errors"

// CollatzConjecture Takes an integer and returns the number of steps it takes to get to 1
func CollatzConjecture(n int) (int, error) {
	steps := 0
	if n < 1 {
		return 0, errors.New("n must be at least 1")
	}
	for n > 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = n*3 + 1
		}
		steps++
	}
	return steps, nil
}

package darts

import "math"

//Score returns the score for a thrown dart
func Score(x float64, y float64) int {
	distance := math.Sqrt(math.Pow(x, 2) + math.Pow(y, 2))
	switch {
	case distance <= 1:
		return 10
	case distance <= 5:
		return 5
	case distance <= 10:
		return 1
	}
	return 0
}

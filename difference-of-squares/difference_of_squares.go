package diffsquares

import "math"

// SquareOfSum find the square of the sum of the first N numbers
func SquareOfSum(n int) (total int) {
	for i := 1; i <= n; i++ {
		total += i
	}
	return int(math.Pow(float64(total), 2))
}

// SumOfSquares returns the sum of the squares of the first n numbers
func SumOfSquares(n int) (total int) {
	for i := 1; i <= n; i++ {
		total += int(math.Pow(float64(i), 2))
	}
	return
}

// Difference returns the difference of the SumOfSquares and SquareOfSum
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}

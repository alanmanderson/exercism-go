// Package triangle determines what type of triangle given the edge lengths
package triangle

import "math"

// Kind denotes the kind of triangle
type Kind string

const (
	// NaT is Not a Triangle
	NaT = "Not a triangle"
	// Equ Equilatera
	Equ = "equilateral"
	// Iso Isoceles
	Iso = "isosceles"
	// Sca Scalene
	Sca = "scalene"
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}
	if a == 0 || b == 0 || c == 0 || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca
}

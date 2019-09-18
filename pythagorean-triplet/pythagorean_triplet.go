package pythagorean

import "math"

// Triplet stuff
type Triplet [3]int

//Range stuff
func Range(min int, max int) []Triplet {
	triplets := make([]Triplet, 0)
	for i := min; i < max; i++ {
		for j := i; j < max; j++ {
			k := math.Sqrt(math.Pow(float64(i), 2) + math.Pow(float64(j), 2))
			if k > float64(max) {
				continue
			}
			if k == float64(int(k)) {
				triplets = append(triplets, Triplet{i, j, int(k)})
			}
		}
	}
	return triplets
}

//Sum stuff
func Sum(sum int) []Triplet {
	out := make([]Triplet, 0)
	triplets := Range(1, sum/2)
	for _, triplet := range triplets {
		if triplet[0]+triplet[1]+triplet[2] == sum {
			out = append(out, triplet)
		}
	}
	return out
}

package strand

import "strings"

var compliments = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

// ToRNA takes in a dna string and returns the rna compliment
func ToRNA(dna string) string {
	var rna strings.Builder
	for _, char := range dna {
		rna.WriteString(compliments[char])
	}
	return rna.String()
}

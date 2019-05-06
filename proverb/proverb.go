// Package proverb provides a function for creating a proverb from a list of words
package proverb

// Proverb will generate a proverb based on a series of nouns
func Proverb(rhyme []string) []string {
	out := []string{}
	ending := ""
	for i := 0; i < len(rhyme); i++ {
		if ending == "" {
			ending = "And all for the want of a " + rhyme[i] + "."
			continue
		}
		out = append(out, "For want of a "+rhyme[i-1]+" the "+rhyme[i]+" was lost.")
	}
	if ending != "" {
		out = append(out, ending)
	}
	return out
}

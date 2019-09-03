package anagram

// Detect stuff
func Detect(in string, anagrams []string) []string {
	out := make([]string, 0)
	for _, anagram := range anagrams {
		if areAnagrams(in, anagram) {
			out = append(out, anagram)
		}
	}
	return out
}

func areAnagrams(in1 string, in2 string) bool {
	if len(in1) != len(in2) {
		return false
	}
	return true
}

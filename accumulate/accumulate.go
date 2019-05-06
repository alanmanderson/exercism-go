package accumulate

type mutator func(string) string

// Accumulate call the passed in function on all strings in the input
func Accumulate(in []string, f mutator) (out []string) {
	out = in[:]
	for index, value := range in {
		out[index] = f(value)
	}
	return
}

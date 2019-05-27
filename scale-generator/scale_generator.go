package scale

import (
	"strings"
)

// Scale returns the scale given a tonic and string of intervals
func Scale(tonic string, intervals string) []string {
	var scale []string
	switch tonic {
	case "C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#":
		scale = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	default:
		scale = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
	}
	if intervals == "" {
		intervals = "mmmmmmmmmmmm"
	}
	out := make([]string, len(intervals))
	out[0] = strings.Title(tonic)

	curIndex := 0
	for i, n := range scale {
		if n == out[0] {
			curIndex = i
			break
		}
	}
	scaleLength := len(scale)
	for i, rune := range intervals[0 : len(intervals)-1] {
		if rune == 'm' {
			curIndex = (curIndex + 1) % scaleLength
		} else if rune == 'M' {
			curIndex = (curIndex + 2) % scaleLength
		} else if rune == 'A' {
			curIndex = (curIndex + 3) % scaleLength
		}
		out[i+1] = scale[curIndex]
	}
	return out
}

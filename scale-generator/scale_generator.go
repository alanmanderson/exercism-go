package scale

import "strings"

var sharps = []string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
var flats = []string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}

var sharpTonics = []string{"C", "a", "G", "D", "A", "E", "B", "F#", "e", "b", "f#", "c#", "g#", "d#"}

var flatTonics = []string{"F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb"}

// Scale returns the scale given a tonic and string of intervals
func Scale(tonic string, intervals string) []string {
	var scale []string
	if indexOf(tonic, sharpTonics) > -1 {
		scale = sharps
	} else if indexOf(tonic, flatTonics) > -1 {
		scale = flats
	}
	startingNote := strings.Title(tonic)
	if intervals == "" {
		intervals = "mmmmmmmmmmmm"
	}
	out := make([]string, len(intervals))
	out[0] = startingNote
	curIndex := indexOf(startingNote, scale)
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

func indexOf(needle string, haystack []string) int {
	for i, n := range haystack {
		if n == needle {
			return i
		}
	}
	return -1
}

package rectangles

import "fmt"

type coordinate [2]int

//Count returns the number of rectangles in an ASCII diagram
func Count(diagram []string) int {
	coordinates := getCorners(diagram)
	for _, c := range coordinates {
		fmt.Printf("%v", c)
	}
	return 0
}

func getCorners(diagram []string) []coordinate {
	coordinates := make([]coordinate, 0)
	for y, row := range diagram {
		for x, r := range row {
			if r == '+' {
				coordinates = append(coordinates, coordinate{x, y})
			}
		}
	}
	return coordinates
}

func getCoordinatesOnAxis(coordinates []coordinate, value int, axis rune) (out []coordinate) {
	var idx int
	switch axis {
	case 'x':
		idx = 0
	case 'y':
		idx = 1
	default:
		panic("axis must be x or y")
	}
	for _, coordinate := range coordinates {
		if coordinate[idx] == value {
			out = append(out, coordinate)
		}
	}
	return
}

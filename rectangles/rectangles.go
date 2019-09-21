package rectangles

import "fmt"

type coordinate [2]int

//Count returns the number of rectangles in an ASCII diagram
func Count(diagram []string) int {
	coordinates := make([]coordinate, 0)
	for y, row := range diagram {
		for x, r := range row {
			if r == '+' {
				coordinates = append(coordinates, coordinate{x, y})
			}
		}
	}
	for _, c := range coordinates {
		fmt.Printf("%v", c)
	}
	return 0
}

package rectangles

import "fmt"

type coordinate [2]int

//Count returns the number of rectangles in an ASCII diagram
func Count(diagram []string) int {
	coordinates := getCorners(diagram)
	xCoordinateMap, yCoordinateMap := getCoordinateMaps(coordinates)
	fmt.Printf("%v", xCoordinateMap)
	fmt.Printf("%v", yCoordinateMap)
	for i, c := range coordinates {
		for _, pair := range coordinates[i:] {
			fmt.Println("map: ", pair)
			fmt.Printf("%v", c)
		}
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

func getCoordinateMaps(coordinates []coordinate) (x map[int][]coordinate, y map[int][]coordinate) {
	for _, c := range coordinates {
		if x[c[0]] == nil {
			x[c[0]] = make([]coordinate, 1)
		}
		if y[c[1]] == nil {
			y[c[0]] = make([]coordinate, 1)
		}
		x[c[0]] = append(x[c[0]], c)
		y[c[1]] = append(y[c[1]], c)
	}
	return
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

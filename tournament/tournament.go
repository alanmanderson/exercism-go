package tournament

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type gameData map[string][5]int

// Tally returns the tally of wins by team
func Tally(r io.Reader, w io.Writer) error {
	data := getReaderData(r)
	games := strings.Split(data, "\n")
	results := make(gameData, len(games))
	for _, game := range games {
		values := strings.Split(game, ";")
		if len(values) != 3 {
			continue
		}
		stat1 := results[values[0]]
		stat2 := results[values[1]]
		switch values[2] {
		case "win":
			updateRow("win", &stat1)
			updateRow("loss", &stat2)
		case "loss":
			updateRow("win", &stat2)
			updateRow("loss", &stat1)
		case "draw":
			updateRow("draw", &stat1)
			updateRow("draw", &stat2)
		default:
			return errors.New("Invalid game result")
		}
		results[values[0]] = stat1
		results[values[1]] = stat2
	}
	table := generateTable(results)
	fmt.Fprintf(w, "%s", table)
	return nil
}

func updateRow(result string, row *[5]int) error {
	switch result {
	case "win":
		row[1]++
		row[4] = row[4] + 3
	case "loss":
		row[2]++
	case "draw":
		row[3]++
		row[4]++
	default:
		return errors.New("Invalid result")
	}
	row[0]++
	return nil
}

func getReaderData(r io.Reader) string {
	var sb strings.Builder
	b := make([]byte, 8)
	for {
		bytes, err := r.Read(b)
		if err == io.EOF {
			break
		}
		sb.WriteString(string(b[0:bytes]))
	}
	return sb.String()
}

func generateTable(data gameData) string {
	var sb strings.Builder
	sortedKeys := getSortedKeys(data)
	fmt.Printf("%v", sortedKeys)

	sb.WriteString(fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P"))
	for _, teamName := range sortedKeys {
		sb.WriteString(fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n", teamName, data[teamName][0], data[teamName][1], data[teamName][3], data[teamName][2], data[teamName][4]))
	}
	return sb.String()
}

func getSortedKeys(data gameData) (out []string) {
	for teamName := range data {
		out = append(out, teamName)
	}
	sortedIndex := len(out)
	for sortedIndex > 0 {
		for pos, teamName := range out {
			if pos == sortedIndex-1 {
				break
			}
			if data[teamName][4] < data[out[pos+1]][4] {
				tmp := out[pos]
				out[pos] = out[pos+1]
				out[pos+1] = tmp
			}
		}
		sortedIndex--
	}
	return
}

func main() {

	fmt.Printf("Hello World")

}

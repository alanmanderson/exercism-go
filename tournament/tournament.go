package tournament

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// Tally returns the tally of wins by team
func Tally(r io.Reader, w io.Writer) error {
	data := getReaderData(r)
	games := strings.Split(data, "\n")
	//var results map[string][]int
	results := make(map[string][5]int, len(games))
	for _, game := range games {
		fmt.Printf(game)
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

func generateTable(data map[string][5]int) string {
	var sb strings.Builder
	//var sortedKeys []string
	//for team, stats := range data {

	//}

	sb.WriteString(fmt.Sprintf("%-35s| MP | W | D | L | P\n", "Team"))
	for team, stats := range data {
		sb.WriteString(fmt.Sprintf("%-35s| %-1d | %-1d | %-1d | %-1d\n", team, stats[0], stats[2], stats[3], stats[4]))
	}
	return sb.String()
}

func main() {

	fmt.Printf("Hello World")

}

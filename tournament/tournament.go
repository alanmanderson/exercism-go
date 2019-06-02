package tournament

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

type gameData map[string]stats

type stats struct {
	matchesPlayed int
	wins          int
	losses        int
	draws         int
	points        int
}

func (s *stats) addWin() {
	s.wins++
	s.matchesPlayed++
	s.points += 3
}

func (s *stats) addLoss() {
	s.losses++
	s.matchesPlayed++
}

func (s *stats) addDraw() {
	s.draws++
	s.matchesPlayed++
	s.points++
}

// Tally returns the tally of wins by team
func Tally(r io.Reader, w io.Writer) error {
	data := getReaderData(r)
	games := strings.Split(data, "\n")
	results := make(gameData, len(games))
	for _, game := range games {
		if len(game) == 0 || game[0] == '#' {
			continue
		}
		values := strings.Split(game, ";")
		if len(values) != 3 {
			return errors.New("Invalid game row")
		}
		stat1 := results[values[0]]
		stat2 := results[values[1]]
		switch values[2] {
		case "win":
			stat1.addWin()
			stat2.addLoss()
		case "loss":
			stat1.addLoss()
			stat2.addWin()
		case "draw":
			stat1.addDraw()
			stat2.addDraw()
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

	sb.WriteString(fmt.Sprintf("%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P"))
	for _, teamName := range sortedKeys {
		sb.WriteString(fmt.Sprintf("%-31s|%3d |%3d |%3d |%3d |%3d\n",
			teamName, data[teamName].matchesPlayed, data[teamName].wins,
			data[teamName].draws, data[teamName].losses, data[teamName].points))
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
			if data[teamName].points == data[out[pos+1]].points && teamName < out[pos+1] {
				continue
			}
			if data[teamName].points <= data[out[pos+1]].points {
				tmp := out[pos]
				out[pos] = out[pos+1]
				out[pos+1] = tmp
			}
		}
		sortedIndex--
	}
	return
}

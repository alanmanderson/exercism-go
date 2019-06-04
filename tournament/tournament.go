package tournament

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"sort"
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
	csvReader := csv.NewReader(r)
	csvReader.Comma = ';'
	csvReader.Comment = '#'
	data, err := csvReader.ReadAll()
	if err != nil {
		return err
	}
	results := make(gameData, len(data))
	for _, game := range data {
		if len(game) != 3 {
			return errors.New("Invalid game row")
		}
		stat1 := results[game[0]]
		stat2 := results[game[1]]
		switch game[2] {
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
		results[game[0]] = stat1
		results[game[1]] = stat2
	}
	generateTable(results, w)
	return nil
}

func generateTable(data gameData, w io.Writer) {
	sortedKeys := getSortedKeys(data)
	fmt.Fprintf(w, "%-31s|%3s |%3s |%3s |%3s |%3s\n", "Team", "MP", "W", "D", "L", "P")
	for _, teamName := range sortedKeys {
		fmt.Fprintf(w, "%-31s|%3d |%3d |%3d |%3d |%3d\n",
			teamName, data[teamName].matchesPlayed, data[teamName].wins,
			data[teamName].draws, data[teamName].losses, data[teamName].points)
	}
}

func getSortedKeys(data gameData) (out []string) {
	for teamName := range data {
		out = append(out, teamName)
	}
	sort.Slice(out, func(i, j int) bool {
		return data[out[i]].points > data[out[j]].points || data[out[i]].points == data[out[j]].points && out[i] < out[j]
	})
	return
}

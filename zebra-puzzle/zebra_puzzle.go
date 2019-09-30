package zebra

//Solution stuff
type Solution struct {
	DrinksWater string
	OwnsZebra   string
}

type house struct {
	order      int
	inhabitant string
	color      string
	pet        string
	beverage   string
	cigarette  string
}

//SolvePuzzle stuff
func SolvePuzzle() Solution {
	houses := [5]house{
		house{1, "Englishman", "red", "", "", ""},
		house{2, "", "", "", "", ""},
		house{3, "", "", "", "", ""},
		house{4, "", "", "", "", ""},
		house{5, "", "", "", "", ""},
	}
	solution := Solution{"", ""}
	for _, house := range houses {
		if house.beverage == "water" {
			solution.DrinksWater = house.inhabitant
		}
		if house.pet == "zebra" {
			solution.OwnsZebra = house.inhabitant
		}
	}
	return solution
}

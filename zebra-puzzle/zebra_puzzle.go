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

var inhabitants = []string{
	"Englishman",
	"Spaniard",
	"Norwegian",
	"Japanese",
	"Ukranian",
}

var cigarettes = []string{
	"Old Gold",
	"Kools",
	"Chesterfields",
	"Lucky Strike",
	"Parliaments",
}

var beverages = []string{
	"Coffee",
	"Tea",
	"Milk",
	"Orange juice",
	"Water",
}

var colors = []string{
	"Red",
	"Green",
	"Ivory",
	"Yellow",
	"Blue",
}

var pets = []string{
	"Dog",
	"Snails",
	"Fox",
	"Horse",
	"Zebra",
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
	boolHouses := [5]boolHouse{
		boolHouse{},
		boolHouse{},
		boolHouse{},
		boolHouse{},
		boolHouse{},
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

func setup(boolHouses *[5]boolHouse) {
	boolHouses.setInhabitant("Englishman")
	boolHouses.setColor("Red")
	for i := 1; i < len(boolhouses); i++ {
		boolHouses[i].negateColor("Red")
		boolHouses[i].negateInhabitant("Englishman")
	}
	boolHouses[0].negatePet("Dog")
	boolHouses[0].negateDrink("Coffee")
	boolHouses[0].negateDrink("Tea")
	boolHouses[1].setInhabitant("Norwegian")
	boolHouses[1].isFirst = true
	boolHouses[1].negatePet("Dog")
	boolHouses[1].negateCigarette("Parliament")
	boolHouses[2].isSecond = true
	boolHouses[2].setColor("Blue")

}

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
	setup(&boolHouses)
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
	boolHouses[0].setInhabitant("Englishman")
	boolHouses[0].setColor("Red")
	negateAttribute(boolHouses, "Color", "Red", 0)
	negateAttribute(boolHouses, "Inhabitant", "Englishman", 0)

	boolHouses[1].setInhabitant("Spaniard")
	boolHouses[1].setPet("Dog")
	negateAttribute(boolHouses, "Pet", "Dog", 1)
	negateAttribute(boolHouses, "Inhabitant", "Spaniard", 1)

	boolHouses[2].setInhabitant("Ukranian")
	boolHouses[2].setDrink("Tea")
	negateAttribute(boolHouses, "Inhabitant", "Ukranian", 2)
	negateAttribute(boolHouses, "Drink", "Tea", 2)

	boolHouses[0].negateDrink("Coffee")
	boolHouses[1].setInhabitant("Norwegian")
	boolHouses[1].setOrder(1)
	boolHouses[1].negateCigarette("Parliament")
	boolHouses[2].setOrder(2)
	boolHouses[2].setColor("Blue")

}

func negateAttribute(boolHouses *[5]boolHouse, attName string, attValue string, exclude int) {
	for i := 0; i < len(boolHouses); i++ {
		if i == exclude {
			continue
		}
		switch attName {
		case "Drink":
			boolHouses[i].negateDrink(attValue)
		case "Inhabitant":
			boolHouses[i].negateInahbitant(attValue)
		case "Pet":
			boolHouses[i].negatePet(attValue)
		case "Cigarette":
			boolHouses[i].negateCigarette(attValue)
		case "Color":
			boolHouses[i].negateColor(attValue)
		}
	}
}

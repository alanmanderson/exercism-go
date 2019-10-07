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
	houses := [5]boolHouse{
		boolHouse{},
		boolHouse{},
		boolHouse{},
		boolHouse{},
		boolHouse{},
	}
	setup(&houses)
	solution := Solution{"", ""}
	for _, house := range houses {
		if house.drinksWater {
			solution.DrinksWater = house.getInhabitant()
		}
		if house.ownsZebra {
			solution.OwnsZebra = house.getInhabitant()
		}
	}
	return solution
}

func setup(houses *[5]boolHouse) {
	houses[0].setInhabitant("Englishman")
	houses[0].setColor("Red")
	negateAttribute(houses, "Color", "Red", 0)
	negateAttribute(houses, "Inhabitant", "Englishman", 0)

	houses[1].setInhabitant("Spaniard")
	houses[1].setPet("Dog")
	negateAttribute(houses, "Pet", "Dog", 1)
	negateAttribute(houses, "Inhabitant", "Spaniard", 1)

	houses[2].setInhabitant("Ukranian")
	houses[2].setDrink("Tea")
	negateAttribute(houses, "Inhabitant", "Ukranian", 2)
	negateAttribute(houses, "Drink", "Tea", 2)

	houses[0].negateDrink("Coffee")
	houses[1].setInhabitant("Norwegian")
	houses[1].setOrder(1)
	houses[1].negateCigarette("Parliament")
	houses[2].setOrder(2)
	houses[2].setColor("Blue")

}

func negateAttribute(houses *[5]boolHouse, attName string, attValue string, exclude int) {
	for i := 0; i < len(houses); i++ {
		if i == exclude {
			continue
		}
		switch attName {
		case "Drink":
			houses[i].negateDrink(attValue)
		case "Inhabitant":
			houses[i].negateInahbitant(attValue)
		case "Pet":
			houses[i].negatePet(attValue)
		case "Cigarette":
			houses[i].negateCigarette(attValue)
		case "Color":
			houses[i].negateColor(attValue)
		}
	}
}

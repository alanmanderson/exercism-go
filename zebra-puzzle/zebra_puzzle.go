package zebra

import "fmt"

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
	for _, h := range houses {
		if h.drinksWater != nil && *h.drinksWater {
			solution.DrinksWater = h.getInhabitant()
		}
		if h.ownsZebra != nil && *h.ownsZebra {
			solution.OwnsZebra = h.getInhabitant()
		}
	}
	return solution
}

func setup(houses *[5]boolHouse) {
	// 2. The Englishman lives in the red house.
	houses[0].setInhabitant(englishman)
	houses[0].setColor(red)
	negateAttribute(houses, color, red, 0)
	negateAttribute(houses, inhabitant, englishman, 0)

	// 3. The Spaniard owns the dog.
	houses[1].setInhabitant(spaniard)
	houses[1].setPet(dog)
	negateAttribute(houses, pet, dog, 1)
	negateAttribute(houses, inhabitant, spaniard, 1)

	// 4. Coffee is drunk in the green house.
	houses[0].negateDrink(coffee)

	// 5. The Ukrainian drinks tea.
	houses[2].setInhabitant(ukrainian)
	houses[2].setDrink(tea)
	negateAttribute(houses, inhabitant, ukrainian, 2)
	negateAttribute(houses, drink, tea, 2)

	// 6. The green house is immediately to the right of the ivory house.

	// 7. The Old Gold smoker owns snails.

	// 8. Kools are smoked in the yellow house.

	// 9. Milk is drunk in the middle house.
	houses[3].negateDrink(milk)
	houses[2].negateOrder("3")

	// 10. The Norwegian lives in the first house.
	houses[3].setInhabitant(norwegian)
	houses[3].setOrder("1")
	negateAttribute(houses, order, "1", 3)
	negateAttribute(houses, inhabitant, norwegian, 3)

	// 11. The man who smokes Chesterfields lives in the house next to the man with the fox.

	// 12. Kools are smoked in the house next to the house where the horse is kept.

	// 13. The Lucky Strike smoker drinks orange juice.
	houses[2].negateDrink(orangeJuice)

	// 14. The Japanese smokes Parliaments.
	houses[4].setInhabitant(japanese)
	houses[4].setCigarette(parliaments)
	negateAttribute(houses, inhabitant, japanese, 4)
	negateAttribute(houses, cigarette, parliaments, 4)

	// 15. The Norwegian lives next to the blue house.

	for _, h := range houses {
		fmt.Printf("%v", h)
	}
}

func negateAttribute(houses *[5]boolHouse, attName string, attValue string, exclude int) {
	for i := 0; i < len(houses); i++ {
		if i == exclude {
			continue
		}
		switch attName {
		case drink:
			houses[i].negateDrink(attValue)
		case inhabitant:
			houses[i].negateInhabitant(attValue)
		case pet:
			houses[i].negatePet(attValue)
		case cigarette:
			houses[i].negateCigarette(attValue)
		case color:
			houses[i].negateColor(attValue)
		}
	}
}

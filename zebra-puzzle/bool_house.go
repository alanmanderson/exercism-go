package zebra

type boolHouse struct {
	isFirst             *bool
	isSecond            *bool
	isThird             *bool
	isFourth            *bool
	isFifth             *bool
	isEnglishman        *bool
	isSpaniard          *bool
	isNorwegian         *bool
	isJapanese          *bool
	isUkranian          *bool
	smokesOldGold       *bool
	smokesKools         *bool
	smokesChesterfields *bool
	smokesLuckyStrike   *bool
	smokesParliaments   *bool
	drinksCoffee        *bool
	drinksTea           *bool
	drinksMilk          *bool
	drinksOrangeJuice   *bool
	drinksWater         *bool
	paintedRed          *bool
	paintedGreen        *bool
	paintedIvory        *bool
	paintedYellow       *bool
	paintedBlue         *bool
	ownsDog             *bool
	ownsSnails          *bool
	ownsFox             *bool
	ownsHorse           *bool
	ownsZebra           *bool
}

func (h *boolHouse) setInhabitant(i string) {
	switch i {
	case "Englishman":
		h.isEnglishman = newTrue()
		h.isSpaniard = newFalse()
		h.isUkranian = newFalse()
		h.isJapanese = newFalse()
		h.isNorwegian = newFalse()
	case "Spaniard":
		h.isEnglishman = newFalse()
		h.isSpaniard = newTrue()
		h.isUkranian = newFalse()
		h.isJapanese = newFalse()
		h.isNorwegian = newFalse()
	case "Ukranian":
		h.isEnglishman = newFalse()
		h.isSpaniard = newFalse()
		h.isUkranian = newTrue()
		h.isJapanese = newFalse()
		h.isNorwegian = newFalse()
	case "Japanese":
		h.isEnglishman = newFalse()
		h.isSpaniard = newFalse()
		h.isUkranian = newFalse()
		h.isJapanese = newTrue()
		h.isNorwegian = newFalse()
	case "Norwegian":
		h.isEnglishman = newFalse()
		h.isSpaniard = newFalse()
		h.isUkranian = newFalse()
		h.isJapanese = newFalse()
		h.isNorwegian = newTrue()
	}
}

func (h *boolHouse) negateInhabitant(i string) {
	switch i {
	case "Englishman":
		h.isEnglishman = newFalse()
	case "Spaniard":
		h.isSpaniard = newFalse()
	case "Ukranian":
		h.isUkranian = newFalse()
	case "Norwegian":
		h.isNorwegian = newFalse()
	case "Japanese":
		h.isJapanese = newFalse()
	}
}

func (h *boolHouse) setOrder(i int) {
	switch i {
	case 1:
		h.isFirst = newTrue()
		h.isSecond = newFalse()
		h.isThird = newFalse()
		h.isFourth = newFalse()
		h.isFifth = newFalse()
	case 2:
		h.isFirst = newFalse()
		h.isSecond = newTrue()
		h.isThird = newFalse()
		h.isFourth = newFalse()
		h.isFifth = newFalse()
	case 3:
		h.isFirst = newFalse()
		h.isSecond = newFalse()
		h.isThird = newTrue()
		h.isFourth = newFalse()
		h.isFifth = newFalse()
	case 4:
		h.isFirst = newFalse()
		h.isSecond = newFalse()
		h.isThird = newFalse()
		h.isFourth = newTrue()
		h.isFifth = newFalse()
	case 5:
		h.isFirst = newFalse()
		h.isSecond = newFalse()
		h.isThird = newFalse()
		h.isFourth = newFalse()
		h.isFifth = newTrue()
	}
}

func (h *boolHouse) negateOrder(i int) {
	switch i {
	case 1:
		h.isFirst = newFalse()
	case 2:
		h.isSecond = newFalse()
	case 3:
		h.isThird = newFalse()
	case 4:
		h.isFourth = newFalse()
	case 5:
		h.isFifth = newFalse()
	}
}

func (h *boolHouse) setDrink(drink string) {
	switch drink {
	case "Milk":
		h.drinksMilk = newTrue()
		h.drinksWater = newFalse()
		h.drinksTea = newFalse()
		h.drinksCoffee = newFalse()
		h.drinksOrangeJuice = newFalse()
	case "Water":
		h.drinksMilk = newFalse()
		h.drinksWater = newTrue()
		h.drinksTea = newFalse()
		h.drinksCoffee = newFalse()
		h.drinksOrangeJuice = newFalse()
	case "Tea":
		h.drinksMilk = newFalse()
		h.drinksWater = newFalse()
		h.drinksTea = newTrue()
		h.drinksCoffee = newFalse()
		h.drinksOrangeJuice = newFalse()
	case "Coffee":
		h.drinksMilk = newFalse()
		h.drinksWater = newFalse()
		h.drinksTea = newFalse()
		h.drinksCoffee = newTrue()
		h.drinksOrangeJuice = newFalse()
	case "Orange Juice":
		h.drinksMilk = newFalse()
		h.drinksWater = newFalse()
		h.drinksTea = newFalse()
		h.drinksCoffee = newFalse()
		h.drinksOrangeJuice = newTrue()
	}
}

func (h *boolHouse) negateDrink(drink string) {
	switch drink {
	case "Milk":
		h.drinksMilk = newFalse()
	case "Water":
		h.drinksWater = newFalse()
	case "Tea":
		h.drinksTea = newFalse()
	case "Coffee":
		h.drinksCoffee = newFalse()
	case "Orange Juice":
		h.drinksOrangeJuice = newFalse()
	}
}

func (h *boolHouse) setCigarette(cigarette string) {
	switch cigarette {
	case "Old Gold":
		h.smokesOldGold = newTrue()
		h.smokesKools = newFalse()
		h.smokesChesterfields = newFalse()
		h.smokesLuckyStrike = newFalse()
		h.smokesParliaments = newFalse()
	case "Kools":
		h.smokesOldGold = newFalse()
		h.smokesKools = newTrue()
		h.smokesChesterfields = newFalse()
		h.smokesLuckyStrike = newFalse()
		h.smokesParliaments = newFalse()
	case "Chesterfields":
		h.smokesOldGold = newFalse()
		h.smokesKools = newFalse()
		h.smokesChesterfields = newTrue()
		h.smokesLuckyStrike = newFalse()
		h.smokesParliaments = newFalse()
	case "Lucky Strike":
		h.smokesOldGold = newFalse()
		h.smokesKools = newFalse()
		h.smokesChesterfields = newFalse()
		h.smokesLuckyStrike = newTrue()
		h.smokesParliaments = newFalse()
	case "Parliaments":
		h.smokesOldGold = newFalse()
		h.smokesKools = newFalse()
		h.smokesChesterfields = newFalse()
		h.smokesLuckyStrike = newFalse()
		h.smokesParliaments = newTrue()
	}
}

func (h *boolHouse) negateCigarette(cigarette string) {
	switch cigarette {
	case "Old Gold":
		h.smokesOldGold = newFalse()
	case "Kools":
		h.smokesKools = newFalse()
	case "Chesterfields":
		h.smokesChesterfields = newFalse()
	case "Lucky Strike":
		h.smokesLuckyStrike = newFalse()
	case "Parliaments":
		h.smokesParliaments = newFalse()
	}
}

func (h *boolHouse) setColor(color string) {
	switch color {
	case "Red":
		h.paintedRed = newTrue()
		h.paintedGreen = newFalse()
		h.paintedIvory = newFalse()
		h.paintedYellow = newFalse()
		h.paintedBlue = newFalse()
	case "Green":
		h.paintedRed = newFalse()
		h.paintedGreen = newTrue()
		h.paintedIvory = newFalse()
		h.paintedYellow = newFalse()
		h.paintedBlue = newFalse()
	case "Ivory":
		h.paintedRed = newFalse()
		h.paintedGreen = newFalse()
		h.paintedIvory = newTrue()
		h.paintedYellow = newFalse()
		h.paintedBlue = newFalse()
	case "Yellow":
		h.paintedRed = newFalse()
		h.paintedGreen = newFalse()
		h.paintedIvory = newFalse()
		h.paintedYellow = newTrue()
		h.paintedBlue = newFalse()
	case "Blue":
		h.paintedRed = newFalse()
		h.paintedGreen = newFalse()
		h.paintedIvory = newFalse()
		h.paintedYellow = newFalse()
		h.paintedBlue = newTrue()
	}
}

func (h *boolHouse) negateColor(color string) {
	switch color {
	case "Red":
		h.paintedRed = newFalse()
	case "Green":
		h.paintedGreen = newFalse()
	case "Blue":
		h.paintedBlue = newFalse()
	case "Ivory":
		h.paintedIvory = newFalse()
	case "Yellow":
		h.paintedYellow = newFalse()
	}
}

func (h *boolHouse) setPet(pet string) {
	h.ownsDog = newFalse()
	h.ownsSnails = newFalse()
	h.ownsFox = newFalse()
	h.ownsHorse = newFalse()
	h.ownsZebra = newFalse()
	switch pet {
	case "Dog":
		h.ownsDog = newTrue()
	case "Snails":
		h.ownsSnails = newTrue()
	case "Fox":
		h.ownsFox = newTrue()
	case "Horse":
		h.ownsHorse = newTrue()
	case "Zebra":
		h.ownsZebra = newTrue()
	}
}

func (h *boolHouse) negatePet(pet string) {
	switch pet {
	case "Dog":
		h.ownsDog = newFalse()
	case "Snails":
		h.ownsSnails = newFalse()
	case "Fox":
		h.ownsFox = newFalse()
	case "Horse":
		h.ownsHorse = newFalse()
	case "Zebra":
		h.ownsZebra = newFalse()
	}
}

func newTrue() *bool {
	b := true
	return &b
}

func newFalse() *bool {
	b := false
	return &b
}

package zebra

const (
	inhabitant    = "Inhabitant"
	color         = "Color"
	pet           = "Pet"
	order         = "Order"
	drink         = "Drink"
	cigarette     = "Cigarette"
	englishman    = "Englishman"
	spaniard      = "Spaniard"
	norwegian     = "Norwegian"
	japanese      = "Japanese"
	ukrainian     = "Ukrainian"
	oldGold       = "Old Gold"
	kools         = "Kools"
	chesterfields = "Chesterfields"
	luckyStrike   = "Lucky Strike"
	parliaments   = "Parliaments"
	coffee        = "Coffee"
	tea           = "Tea"
	milk          = "Milk"
	orangeJuice   = "Orange Juice"
	water         = "Water"
	red           = "Red"
	green         = "Green"
	ivory         = "Ivory"
	yellow        = "Yellow"
	blue          = "Blue"
	dog           = "Dog"
	snails        = "Snails"
	fox           = "Fox"
	horse         = "Horse"
	zebra         = "Zebra"
)

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
	isUkrainian         *bool
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

func (h boolHouse) getInhabitant() string {
	switch {
	case *h.isEnglishman:
		return englishman
	case *h.isSpaniard:
		return spaniard
	case *h.isUkrainian:
		return ukrainian
	case *h.isJapanese:
		return japanese
	case *h.isNorwegian:
		return norwegian
	}
	panic("Invalid inhabitant")
}

func (h *boolHouse) setInhabitant(i string) {
	switch i {
	case englishman:
		h.isEnglishman = newTrue()
		h.isSpaniard = newFalse()
		h.isUkrainian = newFalse()
		h.isJapanese = newFalse()
		h.isNorwegian = newFalse()
	case spaniard:
		h.isEnglishman = newFalse()
		h.isSpaniard = newTrue()
		h.isUkrainian = newFalse()
		h.isJapanese = newFalse()
		h.isNorwegian = newFalse()
	case ukrainian:
		h.isEnglishman = newFalse()
		h.isSpaniard = newFalse()
		h.isUkrainian = newTrue()
		h.isJapanese = newFalse()
		h.isNorwegian = newFalse()
	case japanese:
		h.isEnglishman = newFalse()
		h.isSpaniard = newFalse()
		h.isUkrainian = newFalse()
		h.isJapanese = newTrue()
		h.isNorwegian = newFalse()
	case norwegian:
		h.isEnglishman = newFalse()
		h.isSpaniard = newFalse()
		h.isUkrainian = newFalse()
		h.isJapanese = newFalse()
		h.isNorwegian = newTrue()
	}
}

func (h *boolHouse) negateInhabitant(i string) {
	switch i {
	case englishman:
		h.isEnglishman = newFalse()
	case spaniard:
		h.isSpaniard = newFalse()
	case ukrainian:
		h.isUkrainian = newFalse()
	case norwegian:
		h.isNorwegian = newFalse()
	case japanese:
		h.isJapanese = newFalse()
	}
}

func (h *boolHouse) setOrder(i string) {
	switch i {
	case "1":
		h.isFirst = newTrue()
		h.isSecond = newFalse()
		h.isThird = newFalse()
		h.isFourth = newFalse()
		h.isFifth = newFalse()
	case "2":
		h.isFirst = newFalse()
		h.isSecond = newTrue()
		h.isThird = newFalse()
		h.isFourth = newFalse()
		h.isFifth = newFalse()
	case "3":
		h.isFirst = newFalse()
		h.isSecond = newFalse()
		h.isThird = newTrue()
		h.isFourth = newFalse()
		h.isFifth = newFalse()
	case "4":
		h.isFirst = newFalse()
		h.isSecond = newFalse()
		h.isThird = newFalse()
		h.isFourth = newTrue()
		h.isFifth = newFalse()
	case "5":
		h.isFirst = newFalse()
		h.isSecond = newFalse()
		h.isThird = newFalse()
		h.isFourth = newFalse()
		h.isFifth = newTrue()
	}
}

func (h *boolHouse) negateOrder(i string) {
	switch i {
	case "1":
		h.isFirst = newFalse()
	case "2":
		h.isSecond = newFalse()
	case "3":
		h.isThird = newFalse()
	case "4":
		h.isFourth = newFalse()
	case "5":
		h.isFifth = newFalse()
	}
}

func (h *boolHouse) setDrink(drink string) {
	switch drink {
	case milk:
		h.drinksMilk = newTrue()
		h.drinksWater = newFalse()
		h.drinksTea = newFalse()
		h.drinksCoffee = newFalse()
		h.drinksOrangeJuice = newFalse()
	case water:
		h.drinksMilk = newFalse()
		h.drinksWater = newTrue()
		h.drinksTea = newFalse()
		h.drinksCoffee = newFalse()
		h.drinksOrangeJuice = newFalse()
	case tea:
		h.drinksMilk = newFalse()
		h.drinksWater = newFalse()
		h.drinksTea = newTrue()
		h.drinksCoffee = newFalse()
		h.drinksOrangeJuice = newFalse()
	case coffee:
		h.drinksMilk = newFalse()
		h.drinksWater = newFalse()
		h.drinksTea = newFalse()
		h.drinksCoffee = newTrue()
		h.drinksOrangeJuice = newFalse()
	case orangeJuice:
		h.drinksMilk = newFalse()
		h.drinksWater = newFalse()
		h.drinksTea = newFalse()
		h.drinksCoffee = newFalse()
		h.drinksOrangeJuice = newTrue()
	}
}

func (h *boolHouse) negateDrink(drink string) {
	switch drink {
	case milk:
		h.drinksMilk = newFalse()
	case water:
		h.drinksWater = newFalse()
	case tea:
		h.drinksTea = newFalse()
	case coffee:
		h.drinksCoffee = newFalse()
	case orangeJuice:
		h.drinksOrangeJuice = newFalse()
	}
}

func (h *boolHouse) setCigarette(cigarette string) {
	switch cigarette {
	case oldGold:
		h.smokesOldGold = newTrue()
		h.smokesKools = newFalse()
		h.smokesChesterfields = newFalse()
		h.smokesLuckyStrike = newFalse()
		h.smokesParliaments = newFalse()
	case kools:
		h.smokesOldGold = newFalse()
		h.smokesKools = newTrue()
		h.smokesChesterfields = newFalse()
		h.smokesLuckyStrike = newFalse()
		h.smokesParliaments = newFalse()
	case chesterfields:
		h.smokesOldGold = newFalse()
		h.smokesKools = newFalse()
		h.smokesChesterfields = newTrue()
		h.smokesLuckyStrike = newFalse()
		h.smokesParliaments = newFalse()
	case luckyStrike:
		h.smokesOldGold = newFalse()
		h.smokesKools = newFalse()
		h.smokesChesterfields = newFalse()
		h.smokesLuckyStrike = newTrue()
		h.smokesParliaments = newFalse()
	case parliaments:
		h.smokesOldGold = newFalse()
		h.smokesKools = newFalse()
		h.smokesChesterfields = newFalse()
		h.smokesLuckyStrike = newFalse()
		h.smokesParliaments = newTrue()
	}
}

func (h *boolHouse) negateCigarette(cigarette string) {
	switch cigarette {
	case oldGold:
		h.smokesOldGold = newFalse()
	case kools:
		h.smokesKools = newFalse()
	case chesterfields:
		h.smokesChesterfields = newFalse()
	case luckyStrike:
		h.smokesLuckyStrike = newFalse()
	case parliaments:
		h.smokesParliaments = newFalse()
	}
}

func (h *boolHouse) setColor(color string) {
	switch color {
	case red:
		h.paintedRed = newTrue()
		h.paintedGreen = newFalse()
		h.paintedIvory = newFalse()
		h.paintedYellow = newFalse()
		h.paintedBlue = newFalse()
	case green:
		h.paintedRed = newFalse()
		h.paintedGreen = newTrue()
		h.paintedIvory = newFalse()
		h.paintedYellow = newFalse()
		h.paintedBlue = newFalse()
	case ivory:
		h.paintedRed = newFalse()
		h.paintedGreen = newFalse()
		h.paintedIvory = newTrue()
		h.paintedYellow = newFalse()
		h.paintedBlue = newFalse()
	case yellow:
		h.paintedRed = newFalse()
		h.paintedGreen = newFalse()
		h.paintedIvory = newFalse()
		h.paintedYellow = newTrue()
		h.paintedBlue = newFalse()
	case blue:
		h.paintedRed = newFalse()
		h.paintedGreen = newFalse()
		h.paintedIvory = newFalse()
		h.paintedYellow = newFalse()
		h.paintedBlue = newTrue()
	}
}

func (h *boolHouse) negateColor(color string) {
	switch color {
	case red:
		h.paintedRed = newFalse()
	case green:
		h.paintedGreen = newFalse()
	case blue:
		h.paintedBlue = newFalse()
	case ivory:
		h.paintedIvory = newFalse()
	case yellow:
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
	case dog:
		h.ownsDog = newTrue()
	case snails:
		h.ownsSnails = newTrue()
	case fox:
		h.ownsFox = newTrue()
	case horse:
		h.ownsHorse = newTrue()
	case zebra:
		h.ownsZebra = newTrue()
	}
}

func (h *boolHouse) negatePet(pet string) {
	switch pet {
	case dog:
		h.ownsDog = newFalse()
	case snails:
		h.ownsSnails = newFalse()
	case fox:
		h.ownsFox = newFalse()
	case horse:
		h.ownsHorse = newFalse()
	case zebra:
		h.ownsZebra = newFalse()
	}
}

func isTrue(b *bool) bool {
	return b != nil && *b
}

func newTrue() *bool {
	b := true
	return &b
}

func newFalse() *bool {
	b := false
	return &b
}

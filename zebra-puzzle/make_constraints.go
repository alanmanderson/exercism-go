package zebra

func makeConstraints() []constraint {
	return []constraint{
		constraint{
			func(houses [5]boolHouse) bool {
				for _, h := range houses {
					bools := []*bool{
						h.isFirst,
						h.isSecond,
						h.isThird,
						h.isFourth,
						h.isFifth,
					}
					if countTrues(bools) > 1 {
						return false
					}
				}
				return true
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				for _, h := range houses {
					bools := []*bool{
						h.isEnglishman,
						h.isSpaniard,
						h.isNorwegian,
						h.isJapanese,
						h.isUkrainian,
					}
					if countTrues(bools) > 1 {
						return false
					}
				}
				return true
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				for _, h := range houses {
					bools := []*bool{
						h.smokesOldGold,
						h.smokesKools,
						h.smokesChesterfields,
						h.smokesLuckyStrike,
						h.smokesParliaments,
					}
					if countTrues(bools) > 1 {
						return false
					}
				}
				return true
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				for _, h := range houses {
					bools := []*bool{
						h.drinksCoffee,
						h.drinksTea,
						h.drinksMilk,
						h.drinksOrangeJuice,
						h.drinksWater,
					}
					if countTrues(bools) > 1 {
						return false
					}
				}
				return true
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				for _, h := range houses {
					bools := []*bool{
						h.paintedRed,
						h.paintedYellow,
						h.paintedIvory,
						h.paintedGreen,
						h.paintedBlue,
					}
					if countTrues(bools) > 1 {
						return false
					}
				}
				return true
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				for _, h := range houses {
					bools := []*bool{
						h.ownsDog,
						h.ownsSnails,
						h.ownsFox,
						h.ownsHorse,
						h.ownsZebra,
					}
					if countTrues(bools) > 1 {
						return false
					}
				}
				return true
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].ownsDog,
					houses[1].ownsDog,
					houses[2].ownsDog,
					houses[3].ownsDog,
					houses[4].ownsDog,
				}
				return countTrues(bools) <= 1
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].ownsFox,
					houses[1].ownsFox,
					houses[2].ownsFox,
					houses[3].ownsFox,
					houses[4].ownsFox,
				}
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].ownsHorse,
					houses[1].ownsHorse,
					houses[2].ownsHorse,
					houses[3].ownsHorse,
					houses[4].ownsHorse,
				}
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].ownsZebra,
					houses[1].ownsZebra,
					houses[2].ownsZebra,
					houses[3].ownsZebra,
					houses[4].ownsZebra,
				}
			},
		},
	}
}

func countTrues(bools []*bool) int {
	trueCount := 0
	for _, b := range bools {
		if isTrue(b) == true {
			trueCount++
		}
	}
	return trueCount
}

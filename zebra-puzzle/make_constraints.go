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
				return countTrues(bools) <= 1
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
				return countTrues(bools) <= 1
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
				return countTrues(bools) <= 1
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].paintedRed,
					houses[1].paintedRed,
					houses[2].paintedRed,
					houses[3].paintedRed,
					houses[4].paintedRed,
				}
				return countTrues(bools) <= 1
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].paintedYellow,
					houses[1].paintedYellow,
					houses[2].paintedYellow,
					houses[3].paintedYellow,
					houses[4].paintedYellow,
				}
				return countTrues(bools) <= 1
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].paintedIvory,
					houses[1].paintedIvory,
					houses[2].paintedIvory,
					houses[3].paintedIvory,
					houses[4].paintedIvory,
				}
				return countTrues(bools) <= 1
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].paintedGreen,
					houses[1].paintedGreen,
					houses[2].paintedGreen,
					houses[3].paintedGreen,
					houses[4].paintedGreen,
				}
				return countTrues(bools) <= 1
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].isFirst,
					houses[1].isFirst,
					houses[2].isFirst,
					houses[3].isFirst,
					houses[4].isFirst,
				}
				return countTrues(bools) <= 1
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].isSecond,
					houses[1].isSecond,
					houses[2].isSecond,
					houses[3].isSecond,
					houses[4].isSecond,
				}
				return countTrues(bools) <= 1
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].isThird,
					houses[1].isThird,
					houses[2].isThird,
					houses[3].isThird,
					houses[4].isThird,
				}
				return countTrues(bools) <= 1
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].isFourth,
					houses[1].isFourth,
					houses[2].isFourth,
					houses[3].isFourth,
					houses[4].isFourth,
				}
				return countTrues(bools) <= 1
			},
		},
		constraint{
			func(houses [5]boolHouse) bool {
				bools := []*bool{
					houses[0].isFifth,
					houses[1].isFifth,
					houses[2].isFifth,
					houses[3].isFifth,
					houses[4].isFifth,
				}
				return countTrues(bools) <= 1
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

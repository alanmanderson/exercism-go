package zebra

func makeConstraints() []constraint {
	return []constraint{
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
					trueCount := countTrues(bools)
					if trueCount > 1 {
						return false
					}
				}
				return true
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

package zebra

type constraint struct {
	satisfied func(houses [5]boolHouse) bool
}

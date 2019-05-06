package romannumerals

import "errors"

// ToRomanNumeral takes in a number and converts it to a roman numeral
func ToRomanNumeral(in int) (out string, err error) {
	if in <= 0 || in > 3000 {
		return "", errors.New("Invalid input.  Must be between 1 - 3000 inclusive")
	}
	ones := in % 10
	tens := (in % 100) / 10
	hundreds := (in % 1000) / 100
	thousands := (in % 10000) / 1000
	out += convertDigit(thousands, "M", "_", "_")
	out += convertDigit(hundreds, "C", "D", "M")
	out += convertDigit(tens, "X", "L", "C")
	out += convertDigit(ones, "I", "V", "X")
	return out, nil
}

func convertDigit(digit int, unit string, intermediate string, superUnit string) string {
	switch digit {
	case 1:
		return unit
	case 2:
		return unit + unit
	case 3:
		return unit + unit + unit
	case 4:
		return unit + intermediate
	case 5:
		return intermediate
	case 6:
		return intermediate + unit
	case 7:
		return intermediate + unit + unit
	case 8:
		return intermediate + unit + unit + unit
	case 9:
		return unit + superUnit
	default:
		return ""
	}
}

package twelve

import (
	"fmt"
	"strings"
)

// Song returns the twelve days of Christmas song
func Song() string {
	var sb strings.Builder
	for i := 1; i <= 12; i++ {
		sb.WriteString(fmt.Sprintln(Verse(i)))
	}
	return sb.String()
}

// Verse returns the text for a given verse
func Verse(n int) string {
	//	var verses = []string{
	//		"On the first day of Christmas my true love gave to me:",
	//		"On the second day of Christmas my true love gave to me:",
	//		"On the third day of Christmas my true love gave to me:",
	//		"On the fouth day of Christmas my true love gave to me:""
	//	}
	daysOfChristmas := []string{
		"Partridge in a Pear Tree",
		"Turtle Doves",
		"French Hens",
		"Calling Birds",
		"Gold Rings",
		"Geese-a-Laying",
		"Swans-a-Swimming",
		"Maids-a-Milking",
		"Ladies Dancing",
		"Lords-a-Leaping",
		"Pipers Piping",
		"Drummers Drumming",
	}
	var sb strings.Builder
	nthDay, _ := getNthDay(n)
	sb.WriteString(fmt.Sprintf("On the %s day of Christmas my true love gave to me: ", nthDay))
	for i := n; i > 1; i-- {
		_, day := getNthDay(i)
		sb.WriteString(fmt.Sprintf("%s %s, ", day, daysOfChristmas[i-1]))
	}
	and := ""
	if n > 1 {
		and = "and "
	}
	sb.WriteString(fmt.Sprintf("%s%s %s.", and, "a", daysOfChristmas[0]))
	return sb.String()
}

func getNthDay(day int) (string, string) {
	switch day {
	case 1:
		return "first", "a"
	case 2:
		return "second", "two"
	case 3:
		return "third", "three"
	case 4:
		return "fourth", "four"
	case 5:
		return "fifth", "five"
	case 6:
		return "sixth", "six"
	case 7:
		return "seventh", "seven"
	case 8:
		return "eighth", "eight"
	case 9:
		return "ninth", "nine"
	case 10:
		return "tenth", "ten"
	case 11:
		return "eleventh", "eleven"
	case 12:
		return "twelfth", "twelve"
	default:
		return "", ""
	}
}

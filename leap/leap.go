// This package allows you to determine if a year is a leap year
package leap

// IsLeapYear returns true if the year passed in is a leap year
func IsLeapYear(year int) bool {
	if year%400 == 0 {
		return true
	}
	if year%100 == 0 {
		return false
	}
	return year%4 == 0
}

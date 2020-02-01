package leap

var b bool

//"IsLeapYear" takes a var year of type int
//checks if "year" is a leap year
// returns true or false
func IsLeapYear(year int) bool {
	if year%4 == 0 {
		b = true
		if year%100 == 0 {
			if year%400 == 0 {
				b = true

			} else {
				b = false
			}
		}
	}

	return b
}

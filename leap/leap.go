package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}

	if year%100 == 0 && year%400 != 0 {
		return false
	}

	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

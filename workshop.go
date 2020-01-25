package workshop

func IsLeapYear(aYear int) bool {
	y := year(aYear)

	return y.isLeap()
}

type year int

func (y year) isLeap() bool {
	if y%400 == 0 {
		return true
	}

	if y%4 == 0 && y%100 != 0 {
		return true
	}

	return false
}

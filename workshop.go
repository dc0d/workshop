package workshop

import (
	"strings"
)

func ToRomanNumeral(n int) (result string) {
	result = toRomanNumeral(n)

	result = strings.Replace(result, "LXXXX", "XC", -1)
	result = strings.Replace(result, "VIIII", "IX", -1)
	result = strings.Replace(result, "DCCCC", "CM", -1)
	result = strings.Replace(result, "IIII", "IV", -1)
	result = strings.Replace(result, "XXXX", "XL", -1)
	result = strings.Replace(result, "CCCC", "CD", -1)

	return
}

func toRomanNumeral(n int) (result string) {
	if roman, ok := decToRoman[n]; ok {
		return roman
	}

	{
		ceiling, diff := romanCeiling(n)
		if ceilingNumeral, diffIsRoman := decToRoman[diff]; diffIsRoman {
			return ceilingNumeral + decToRoman[ceiling]
		}
	}

	floor := romanFloor(n)
	divisor, remainder := div(n, floor)
	result = strings.Repeat(decToRoman[floor], divisor)
	if remainder > 0 {
		result += toRomanNumeral(remainder)
	}

	return
}

func div(n, roman int) (q, r int) {
	q = n / roman
	r = n % roman
	return
}

func romanCeiling(n int) (ceiling, diffToCeiling int) {
	switch {
	case n <= 1:
		ceiling, diffToCeiling = 1, 1-n
	case n <= 5:
		ceiling, diffToCeiling = 5, 5-n
	case n <= 10:
		ceiling, diffToCeiling = 10, 10-n
	case n <= 50:
		ceiling, diffToCeiling = 50, 50-n
	case n <= 100:
		ceiling, diffToCeiling = 100, 100-n
	case n <= 500:
		ceiling, diffToCeiling = 500, 500-n
	default:
		ceiling, diffToCeiling = 1000, 1000-n
	}

	return
}

func romanFloor(n int) (roman int) {
	switch {
	case n > 1000:
		return 1000
	case n > 500:
		return 500
	case n > 100:
		return 100
	case n > 50:
		return 50
	case n > 10:
		return 10
	case n > 5:
		return 5
	}
	return 1
}

var decToRoman = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	10:   "X",
	5:    "V",
	1:    "I",
}

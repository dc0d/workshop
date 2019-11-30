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
	if roman, ok := toRomanNumbers[n]; ok {
		return roman
	}

	roman := biggest(n)
	q, r := div(n, roman)
	result += strings.Repeat(toRomanNumbers[roman], q)

	if r > 0 {
		result += toRomanNumeral(r)
	}

	return
}

func div(n, roman int) (q, r int) {
	q = n / roman
	r = n % roman
	return
}

func biggest(n int) (roman int) {
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

var toRomanNumbers = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	10:   "X",
	5:    "V",
	1:    "I",
}

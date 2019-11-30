package workshop_test

import (
	"fmt"
	"testing"

	"github.com/dc0d/workshop"

	assert "github.com/stretchr/testify/require"
)

func Test_draft(t *testing.T) {
	var (
		expectations = []expectation{
			{1, "I"},
			{5, "V"},
			{10, "X"},
			{50, "L"},
			{100, "C"},
			{500, "D"},
			{1000, "M"},
			{7, "VII"},
			{14, "XIV"},
			{15, "XV"},
			{99, "XCIX"},
			{2006, "MMVI"},
			{1944, "MCMXLIV"},
			{3497, "MMMCDXCVII"},
			{1999, "MCMXCIX"}, // also MIM ?
		}
	)

	for _, exp := range expectations {
		var (
			n             = exp.decimal
			expectedRoman = exp.expectedRoman
		)

		t.Run(fmt.Sprintf("converting %d to roman numeral", n), func(t *testing.T) {
			var (
				assert = assert.New(t)
			)

			romanNumber := workshop.ToRomanNumeral(n)
			assert.Equal(expectedRoman, romanNumber)
		})
	}
}

type expectation struct {
	decimal       int
	expectedRoman string
}

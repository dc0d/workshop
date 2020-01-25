package workshop_test

import (
	"testing"

	"github.com/dc0d/workshop"

	assert "github.com/stretchr/testify/require"
)

func Test_a_typical_common_year(t *testing.T) {
	var (
		year       = 2001
		isLeapYear = false

		assert = assert.New(t)
	)

	assert.Equal(isLeapYear, workshop.IsLeapYear(year))
}

func Test_a_typical_leap_year(t *testing.T) {
	var (
		year       = 1996
		isLeapYear = true

		assert = assert.New(t)
	)

	assert.Equal(isLeapYear, workshop.IsLeapYear(year))
}

func Test_a_atypical_common_year(t *testing.T) {
	var (
		year       = 1900
		isLeapYear = false

		assert = assert.New(t)
	)

	assert.Equal(isLeapYear, workshop.IsLeapYear(year))
}

func Test_a_atypical_leap_year(t *testing.T) {
	var (
		year       = 2000
		isLeapYear = true

		assert = assert.New(t)
	)

	assert.Equal(isLeapYear, workshop.IsLeapYear(year))
}

package primefactors_test

import (
	"fmt"
	"testing"

	"github.com/dc0d/primefactors"

	assert "github.com/stretchr/testify/require"
)

func Test_generate_prime_factors(t *testing.T) {
	var (
		expectations = []primeFactorsExpectation{
			{1, nil},
			{2, []int{2}},
			{3, []int{3}},
			{4, []int{2, 2}},
			{5, []int{5}},       // already grean; does it help? or should be removed?
			{6, []int{2, 3}},    // already grean; does it help? or should be removed?
			{7, []int{7}},       // already grean; does it help? or should be removed?
			{8, []int{2, 2, 2}}, // already grean; does it help? or should be removed?
			{9, []int{3, 3}},
			{10, []int{2, 5}},    // already grean; does it help? or should be removed?
			{11, []int{11}},      // already grean; does it help? or should be removed?
			{12, []int{2, 2, 3}}, // already grean; does it help? or should be removed?
			{2 * 3 * 5 * 7 * 11 * 13 * 17 * 19 * 23 * 29, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}}, // first ten prime numbers
		}
	)

	for _, exp := range expectations {
		exp := exp
		t.Run(fmt.Sprintf("generating prime factors for number %d", exp.Num), func(t *testing.T) {
			var (
				assert               = assert.New(t)
				number               = exp.Num
				expectedPrimeFactors = exp.PrimeFactors
			)

			primeFactors := primefactors.Generate(number)
			assert.EqualValues(expectedPrimeFactors, primeFactors)
		})
	}
}

type primeFactorsExpectation struct {
	Num          int
	PrimeFactors []int
}

var _ = assert.New(nil)

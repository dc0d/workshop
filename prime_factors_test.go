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

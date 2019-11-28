package workshop_test

import (
	"fmt"
	"testing"

	"github.com/dc0d/workshop"

	assert "github.com/stretchr/testify/require"
)

func Test_calculate_fibonacci_sequence(t *testing.T) {
	var (
		expectations = []fibExpectation{
			{1, 1},
			{2, 1},
			{3, 2},
			{4, 3},
			{10, 55},
			{12, 144},
			{20, 6765},
		}
	)

	for _, exp := range expectations {
		var (
			n           = exp.n
			expectedFib = exp.fib
		)

		t.Run(fmt.Sprintf("finding %dth fibonacci number", n), func(t *testing.T) {
			var (
				assert = assert.New(t)
			)

			fib := workshop.Fib(n)
			assert.Equal(expectedFib, fib)
		})
	}
}

type fibExpectation struct {
	n   int
	fib int
}

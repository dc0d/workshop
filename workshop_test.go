package workshop_test

import (
	"fmt"
	"testing"

	"github.com/dc0d/workshop"

	"github.com/stretchr/testify/require"
)

func Test_calculate_fibonacci_sequence(t *testing.T) {
	type (
		expectation struct {
			input          int
			expectedOutput int
		}
	)

	var (
		expectations = []expectation{
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
			input          = exp.input
			expectedOutput = exp.expectedOutput
		)

		t.Run(fmt.Sprintf("finding %dth fibonacci number", input), func(t *testing.T) {
			output := workshop.Fib(input)

			require.Equal(t, expectedOutput, output)
		})
	}
}

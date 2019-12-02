package workshop_test

import (
	"fmt"
	"testing"

	"github.com/dc0d/workshop"

	assert "github.com/stretchr/testify/require"
)

func Test_calculate_fibonacci_sequence(t *testing.T) {
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
			var (
				assert = assert.New(t)
			)

			output := workshop.Fib(input)
			assert.Equal(expectedOutput, output)
		})
	}
}

type expectation struct {
	input          int
	expectedOutput int
}

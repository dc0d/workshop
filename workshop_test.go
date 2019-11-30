package workshop_test

import (
	"fmt"
	"testing"

	"github.com/dc0d/workshop"

	assert "github.com/stretchr/testify/require"
)

func Test_add(t *testing.T) {
	var (
		expectations = []expectation{
			{"", "0"},
			{"1", "1"},
			{"1.1,2.2", "3.3"},
			{"1\n2,3", "6"},
			{"175.2,\n35", "Number expected but '\n' found at position 6"},
			{"1,3,", "Number expected but EOF found"},
			{"//;\n1;2", "3"},
			{"//|\n1|2|3", "6"},
			{"//sep\n2sep3", "5"},
			{"//|\n1|2,3", "'|' expected but ',' found at position 3"},
			{"-1,2", "Negative not allowed : -1"},
			{"2,-4,-5", "Negative not allowed : -4, -5"},
			// {"-1,,2", "Negative not allowed : -1\nNumber expected but ',' found at position 3"},
		}
	)

	for _, exp := range expectations {
		var (
			input          = exp.input
			expectedOutput = exp.expectedOutput
		)

		t.Run(fmt.Sprintf("test input %s", input), func(t *testing.T) {
			var (
				assert = assert.New(t)
			)

			result := workshop.Add(input)
			assert.Equal(expectedOutput, result)
		})
	}
}

type expectation struct {
	input          string
	expectedOutput string
}

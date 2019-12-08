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
			{[]string{}, []string{}},
			{[]string{"Alpha 2", "Alpha 1"}, []string{"Alpha 1", "Alpha 2"}},
			{[]string{"Alpha 1", "Alpha 2"}, []string{"Alpha 1", "Alpha 2"}},
			{[]string{"Alpha 100", "Alpha 2A-8000"}, []string{"Alpha 2A-8000", "Alpha 100"}},
			{[]string{"Alpha 100", "Alpha 2"}, []string{"Alpha 2", "Alpha 100"}},
			{
				[]string{
					"Alpha 100",
					"Alpha 2",
					"Alpha 200",
					"Alpha 2A",
					"Alpha 2A-8000",
					"Alpha 2A-900",
				},
				[]string{
					"Alpha 2",
					"Alpha 2A",
					"Alpha 2A-900",
					"Alpha 2A-8000",
					"Alpha 100",
					"Alpha 200",
				},
			},
		}
	)

	for i, exp := range expectations {
		var (
			input          = exp.input
			expectedOutput = exp.expectedOutput
		)

		t.Run(fmt.Sprintf("sort case %v", i), func(t *testing.T) {
			var (
				assert = assert.New(t)
			)

			workshop.Sort(input...)
			output := input
			assert.EqualValues(expectedOutput, output)
		})
	}
}

type expectation struct {
	input          []string
	expectedOutput []string
}

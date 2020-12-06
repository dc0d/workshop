package workshop_test

import (
	"testing"

	"github.com/dc0d/workshop"
	"github.com/stretchr/testify/assert"
)

func TestUpdateQuality(t *testing.T) {
	type expectation struct {
		inputItem      workshop.Item
		expectedOutput workshop.Item
	}

	var (
		expectations = []expectation{
			{
				inputItem:      workshop.Item{Name: "", SellIn: 0, Quality: 0},
				expectedOutput: workshop.Item{Name: "", SellIn: -1, Quality: 0},
			},
			{
				inputItem:      workshop.Item{Name: "NonExisting", SellIn: 0, Quality: 0},
				expectedOutput: workshop.Item{Name: "NonExisting", SellIn: -1, Quality: 0},
			},
			{
				inputItem:      workshop.Item{Name: "", SellIn: 0, Quality: 1},
				expectedOutput: workshop.Item{Name: "", SellIn: -1, Quality: 0},
			},
		}
	)

	for _, exp := range expectations {
		var (
			inputItem      = exp.inputItem
			expectedOutput = exp.expectedOutput
		)

		input := []*workshop.Item{&inputItem}
		workshop.UpdateQuality(input)

		actualOutput := inputItem
		assert.Equal(t, expectedOutput, actualOutput)
	}
}

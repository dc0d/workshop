package workshop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateQuality(t *testing.T) {
	type testCase struct {
		inputItem      Item
		expectedOutput Item
	}

	var (
		testCases = []testCase{
			{
				inputItem:      Item{name: "", sellIn: 0, quality: 0},
				expectedOutput: Item{name: "", sellIn: -1, quality: 0},
			},
			{
				inputItem:      Item{name: "nonexistent", sellIn: 0, quality: 0},
				expectedOutput: Item{name: "nonexistent", sellIn: -1, quality: 0},
			},
			{
				inputItem:      Item{name: "", sellIn: 0, quality: 1},
				expectedOutput: Item{name: "", sellIn: -1, quality: 0},
			},
		}
	)

	for _, testCase := range testCases {
		var (
			inputItem      = testCase.inputItem
			expectedOutput = testCase.expectedOutput
		)

		input := []*Item{&inputItem}
		UpdateQuality(input)

		actualOutput := inputItem
		assert.Equal(t, expectedOutput, actualOutput)
	}
}

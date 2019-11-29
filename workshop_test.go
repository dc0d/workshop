package workshop_test

import (
	"fmt"
	"testing"

	"github.com/dc0d/workshop"

	assert "github.com/stretchr/testify/require"
)

// The Golden Master (Technique)
func Test_cover_paths_in_update_quality(t *testing.T) {
	var (
		expectations = []expectation{
			{
				Input:          workshop.Item{Name: "", SellIn: 0, Quality: 0},
				ExpectedOutput: workshop.Item{Name: "", SellIn: -1, Quality: 0},
			},
			{
				Input:          workshop.Item{Name: "NonExisting", SellIn: 0, Quality: 0},
				ExpectedOutput: workshop.Item{Name: "NonExisting", SellIn: -1, Quality: 0},
			},
			{
				Input:          workshop.Item{Name: "", SellIn: 0, Quality: 1},
				ExpectedOutput: workshop.Item{Name: "", SellIn: -1, Quality: 0},
			},
			{
				Input:          workshop.Item{Name: "", SellIn: -1, Quality: 2},
				ExpectedOutput: workshop.Item{Name: "", SellIn: -2, Quality: 0},
			},
			{
				Input:          workshop.Item{Name: "Aged Brie", SellIn: 0, Quality: 0},
				ExpectedOutput: workshop.Item{Name: "Aged Brie", SellIn: -1, Quality: 2},
			},
			{
				Input:          workshop.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 0},
				ExpectedOutput: workshop.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 9, Quality: 2},
			},
			{
				Input:          workshop.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 5, Quality: 0},
				ExpectedOutput: workshop.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 4, Quality: 3},
			},
			{
				Input:          workshop.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -1, Quality: 0},
				ExpectedOutput: workshop.Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: -2, Quality: 0},
			},
		}
	)

	for _, exp := range expectations {
		var (
			input          = exp.Input
			expectedOutput = exp.ExpectedOutput
		)

		t.Run(fmt.Sprintf("call update quality for %v", input), func(t *testing.T) {
			var (
				assert = assert.New(t)
				input  = []*workshop.Item{&input}
			)

			workshop.UpdateQuality(input)
			output := *input[0]
			assert.Equal(expectedOutput, output)
		})
	}
}

type expectation struct {
	Input          workshop.Item
	ExpectedOutput workshop.Item
}

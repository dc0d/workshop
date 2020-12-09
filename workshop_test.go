package workshop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateQuality(t *testing.T) {
	inputItem := Item{name: "", sellIn: 0, quality: 0}
	expectedOutput := Item{name: "", sellIn: -1, quality: 0}

	input := []*Item{&inputItem}
	UpdateQuality(input)

	actualOutput := inputItem
	assert.Equal(t, expectedOutput, actualOutput)
}

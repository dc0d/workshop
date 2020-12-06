package workshop_test

import (
	"testing"

	"github.com/dc0d/workshop"
	"github.com/stretchr/testify/assert"
)

func TestUpdateQuality(t *testing.T) {
	inputItem := workshop.Item{Name: "", SellIn: 0, Quality: 0}
	expectedOutput := workshop.Item{Name: "", SellIn: -1, Quality: 0}

	input := []*workshop.Item{&inputItem}
	workshop.UpdateQuality(input)

	actualOutput := inputItem
	assert.Equal(t, expectedOutput, actualOutput)
}

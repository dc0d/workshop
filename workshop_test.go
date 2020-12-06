package workshop_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/dc0d/workshop"

	"github.com/stretchr/testify/assert"
)

func TestUpdateQuality(t *testing.T) {
	type testCase struct {
		inputItem workshop.Item
	}

	var (
		expectations = []testCase{
			{
				inputItem: workshop.Item{Name: "", SellIn: 0, Quality: 0},
			},
			{
				inputItem: workshop.Item{Name: "NonExisting", SellIn: 0, Quality: 0},
			},
			{
				inputItem: workshop.Item{Name: "", SellIn: 0, Quality: 1},
			},
		}
	)

	for index, exp := range expectations {
		var (
			inputItem      = exp.inputItem
			expectedOutput workshop.Item
			testCaseName   = filepath.Join("test", "fixture", fmt.Sprintf("%v-%v.json", t.Name(), index))
		)

		fromJSON(readFile(testCaseName), &expectedOutput)

		input := []*workshop.Item{&inputItem}
		workshop.UpdateQuality(input)

		actualOutput := inputItem
		assert.Equal(t, expectedOutput, actualOutput)
	}
}

func readFile(p string) []byte {
	content, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}
	return content
}

func fromJSON(js []byte, ptrVal interface{}) {
	err := json.Unmarshal(js, ptrVal)
	if err != nil {
		panic(err)
	}
}

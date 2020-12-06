package workshop_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dc0d/workshop"
	"github.com/dc0d/workshop/test/support"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			{
				inputItem: workshop.Item{Name: "", SellIn: -1, Quality: 2},
			},
		}
	)

	for index, exp := range expectations {
		var (
			inputItem      = exp.inputItem
			expectedOutput workshop.Item
			fixtureFile    = filepath.Join("test", "fixture", fmt.Sprintf("%v-%v.json", t.Name(), index))
		)

		expectedOutput = getExpectedOutput(t, fixtureFile, inputItem)

		input := []*workshop.Item{&inputItem}
		workshop.UpdateQuality(input)

		actualOutput := inputItem
		assert.Equal(t, expectedOutput, actualOutput)
	}
}

func getExpectedOutput(t *testing.T, fixtureFile string, item workshop.Item) (expectedOutput workshop.Item) {
	js, err := readFile(fixtureFile)
	if err != nil {
		handled := false

		if errors.Is(err, os.ErrNotExist) {
			if *support.UpdateGoldenMaster {
				handleGoldenMaster(t, fixtureFile, item)
				handled = true
			}
		}

		if !handled {
			panic(err)
		}
	}
	fromJSON(js, &expectedOutput)

	return
}

func handleGoldenMaster(t *testing.T, fixtureFile string, item workshop.Item) {
	input := []*workshop.Item{&item}
	workshop.UpdateQuality(input)

	actualOutput := item
	js := toJSON(actualOutput)

	if err := ioutil.WriteFile(fixtureFile, js, 0644); err != nil {
		panic(err)
	}
	require.FailNow(t, "Golden Master is written, run the test again without -update flag")
}

func readFile(p string) ([]byte, error) { return ioutil.ReadFile(p) }

func fromJSON(js []byte, ptrVal interface{}) {
	err := json.Unmarshal(js, ptrVal)
	if err != nil {
		panic(err)
	}
}

func toJSON(val interface{}) []byte {
	js, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	return js
}

package workshop

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/dc0d/workshop/test/support"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateQuality(t *testing.T) {
	var (
		testCases []Item
	)

	for _, name := range []string{"", "nonexistent", "Aged Brie", "Backstage passes to a TAFKAL80ETC concert", "Sulfuras, Hand of Ragnaros"} {
		for _, quality := range []int{0, 1, 49, 50} {
			for _, sellIn := range []int{-1, 0, 5, 6, 10, 11} {
				testCases = append(testCases, Item{name: name, sellIn: sellIn, quality: quality})
			}
		}
	}

	for index, testCase := range testCases {
		var (
			inputItem      = testCase
			expectedOutput Item
			fixtureFile    = filepath.Join("test", "fixture", fmt.Sprintf("%v-%v.json", t.Name(), index))
		)

		expectedOutput = getExpectedOutput(t, fixtureFile, inputItem)

		input := []*Item{&inputItem}
		UpdateQuality(input)

		actualOutput := inputItem
		assert.Equal(t, expectedOutput, actualOutput)
	}
}

func getExpectedOutput(t *testing.T, fixtureFile string, item Item) (expectedOutput Item) {
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
			require.FailNow(t, err.Error())
		}
	}

	var output TestItem
	fromJSON(js, &output)
	expectedOutput = output.to()

	return
}

func handleGoldenMaster(t *testing.T, fixtureFile string, item Item) {
	input := []*Item{&item}
	UpdateQuality(input)

	var actualOutput TestItem
	actualOutput.from(&item)
	js := toJSON(&actualOutput)

	if err := ioutil.WriteFile(fixtureFile, js, 0644); err != nil {
		panic(err)
	}
	require.FailNow(t, "Golden Master is written, run the test again without -update flag")
}

func readFile(p string) ([]byte, error) { return ioutil.ReadFile(p) }

type TestItem struct {
	Name            string
	SellIn, Quality int
}

func (ti *TestItem) from(item *Item) {
	ti.Name = item.name
	ti.SellIn = item.sellIn
	ti.Quality = item.quality
}

func (ti *TestItem) to() Item {
	return Item{
		name:    ti.Name,
		sellIn:  ti.SellIn,
		quality: ti.Quality,
	}
}

func fromJSON(js []byte, ptrVal *TestItem) {
	err := json.Unmarshal(js, ptrVal)
	if err != nil {
		panic(err)
	}
}

func toJSON(val *TestItem) []byte {
	js, err := json.Marshal(val)
	if err != nil {
		panic(err)
	}
	return js
}

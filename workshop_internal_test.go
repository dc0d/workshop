package workshop

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func Test_sortable(t *testing.T) {
	var (
		expectations = []struct {
			input          []string
			expectedOutput []string
		}{
			{
				input:          []string{"Alpha 100", "Alpha 1"},
				expectedOutput: []string{"Alpha 1", "Alpha 100"},
			},
			{
				input:          []string{"Alpha 100", "Alpha 2A-8000"},
				expectedOutput: []string{"Alpha 2A-8000", "Alpha 100"},
			},
		}
	)

	for _, exp := range expectations {
		var (
			input          = exp.input
			expectedOutput = exp.expectedOutput
		)

		t.Run(fmt.Sprintf("for input %v", input), func(t *testing.T) {
			var (
				assert = assert.New(t)
			)

			st := newSortable(input...)
			sort.Sort(st)
			output := st.input
			assert.EqualValues(expectedOutput, output)
		})
	}
}

func Test_partition(t *testing.T) {
	var (
		expectations = []struct {
			input          string
			expectedOutput []segment
		}{
			{
				"Alpha 100",
				[]segment{
					segment{
						str: newStr("Alpha "),
					},
					segment{
						num: newNum(100),
					},
				},
			},
			{
				"Alpha100",
				[]segment{
					segment{
						str: newStr("Alpha"),
					},
					segment{
						num: newNum(100),
					},
				},
			},
			{
				"Alpha 2A-8000",
				[]segment{
					segment{
						str: newStr("Alpha "),
					},
					segment{
						num: newNum(2),
					},
					segment{
						str: newStr("A-"),
					},
					segment{
						num: newNum(8000),
					},
				},
			},
		}
	)

	for _, exp := range expectations {
		var (
			input          = exp.input
			expectedOutput = toJSON(exp.expectedOutput)
		)

		t.Run(fmt.Sprintf("for input %v", input), func(t *testing.T) {
			var (
				assert = assert.New(t)
			)

			output := partition(input)
			assert.EqualValues(expectedOutput, toJSON(output))
		})
	}
}

func Test_less_segment(t *testing.T) {
	var (
		expectations = []struct {
			input1         segment
			input2         segment
			expectedOutput bool
		}{
			{
				input1: segment{
					str: newStr("A"),
				},
				input2: segment{
					str: newStr("B"),
				},
				expectedOutput: true,
			},
			{
				input1: segment{
					num: newNum(10),
				},
				input2: segment{
					num: newNum(20),
				},
				expectedOutput: true,
			},
			{
				input1: segment{
					num: newNum(10),
				},
				input2: segment{
					str: newStr("B"),
				},
				expectedOutput: true,
			},
			{
				input1: segment{
					str: newStr("B"),
				},
				input2: segment{
					num: newNum(10),
				},
				expectedOutput: false,
			},
		}
	)

	for _, exp := range expectations {
		var (
			input1         = exp.input1
			input2         = exp.input2
			expectedOutput = exp.expectedOutput
		)

		t.Run(fmt.Sprintf("compare segments %v and %v", toJSON(input1), toJSON(input2)), func(t *testing.T) {
			var (
				assert = assert.New(t)
			)

			output := lessSegment(input1, input2)
			assert.EqualValues(expectedOutput, output)
		})
	}
}

func toJSON(v interface{}) string {
	if seg, ok := v.(segment); ok {
		var tseg testSegment
		tseg.Str = seg.str
		tseg.Num = seg.num
		v = tseg
	}
	js, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(js)
}

type testSegment struct {
	Str *string `json:"str,omitempty"`
	Num *int    `json:"num,omitempty"`
}

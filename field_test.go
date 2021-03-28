package workshop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Field_validation(t *testing.T) {
	type testCase struct {
		field   Field
		isValid bool
	}

	var testCases []testCase
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			testCases = append(testCases, testCase{field: Field{X: x, Y: y}, isValid: true})
		}
	}

	testCases = append(testCases,
		testCase{field: Field{X: -1, Y: 0}, isValid: false},
		testCase{field: Field{X: 3, Y: 0}, isValid: false},
		testCase{field: Field{X: 0, Y: 3}, isValid: false},
		testCase{field: Field{X: 0, Y: -1}, isValid: false})

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			require.Equal(t, tc.isValid, tc.field.IsValid())
		})
	}
}

package workshop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Player_validation(t *testing.T) {
	type testCase struct {
		player  Player
		isValid bool
	}

	testCases := []testCase{
		{
			player:  PlayerX,
			isValid: true,
		},
		{
			player:  PlayerO,
			isValid: true,
		},
		{
			player:  -1,
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			require.Equal(t, tc.isValid, tc.player.IsValid())
		})
	}
}

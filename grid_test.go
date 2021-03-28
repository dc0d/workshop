package workshop

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Grid_should_return_error_if_field_is_invalid(t *testing.T) {
	field := Field{-1, -1}
	player := PlayerX
	grid := NewGrid()

	require.Equal(t, ErrFieldInvalid, grid.Take(player, field))
}

func Test_Grid_should_return_error_if_player_is_invalid(t *testing.T) {
	field := Field{1, 1}
	var player Player = -1
	grid := NewGrid()

	require.Equal(t, ErrPlayerInvalid, grid.Take(player, field))
}

func Test_Grid_a_player_can_take_a_field_if_not_already_taken(t *testing.T) {
	field := Field{0, 0}
	player := PlayerX
	grid := NewGrid()

	require.False(t, grid.isTaken(field))

	_ = grid.Take(player, field)
	require.True(t, grid.isTaken(field))
}

func Test_Grid_should_return_error_if_field_is_already_taken(t *testing.T) {
	field := Field{0, 0}
	player := PlayerX
	grid := NewGrid()

	require.NoError(t, grid.Take(player, field))
	require.Equal(t, ErrFieldTaken, grid.Take(player, field))
}

func Test_Grid_check_if_all_fields_in_a_column_are_taken_by_a_player(t *testing.T) {
	type testCase struct {
		player Player
		x      int
	}

	testCases := []testCase{
		{PlayerX, 0},
		{PlayerO, 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			g := NewGrid()
			for y := 0; y < 3; y++ {
				err := g.Take(tc.player, Field{tc.x, y})
				require.NoError(t, err)
			}

			require.Equal(t, tc.player, g.Winner())
		})
	}
}

func Test_Grid_check_if_all_fields_in_a_row_are_taken_by_a_player(t *testing.T) {
	type testCase struct {
		player Player
		y      int
	}

	testCases := []testCase{
		{PlayerX, 0},
		{PlayerO, 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%+v", tc), func(t *testing.T) {
			g := NewGrid()
			for x := 0; x < 3; x++ {
				err := g.Take(tc.player, Field{x, tc.y})
				require.NoError(t, err)
			}

			require.Equal(t, tc.player, g.Winner())
		})
	}
}

func Test_Grid_check_if_all_fields_in_a_diagonal_are_taken_by_a_player(t *testing.T) {
	t.Run(`diagonal 1 X`, func(t *testing.T) {
		player := PlayerX
		g := NewGrid()

		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if x != y {
					continue
				}
				err := g.Take(player, Field{x, y})
				require.NoError(t, err)
			}
		}

		require.Equal(t, player, g.Winner())
	})

	t.Run(`diagonal 1 O`, func(t *testing.T) {
		player := PlayerO
		g := NewGrid()

		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if x != y {
					continue
				}
				err := g.Take(player, Field{x, y})
				require.NoError(t, err)
			}
		}

		require.Equal(t, player, g.Winner())
	})

	t.Run(`diagonal 2 X`, func(t *testing.T) {
		player := PlayerX
		g := NewGrid()

		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if x != y {
					continue
				}
				err := g.Take(player, Field{x, 2 - y})
				require.NoError(t, err)
			}
		}

		require.Equal(t, player, g.Winner())
	})

	t.Run(`diagonal 2 O`, func(t *testing.T) {
		player := PlayerO
		g := NewGrid()

		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if x != y {
					continue
				}
				err := g.Take(player, Field{x, 2 - y})
				require.NoError(t, err)
			}
		}

		require.Equal(t, player, g.Winner())
	})
}

func Test_Grid_check_if_all_fields_are_taken(t *testing.T) {
	t.Run(`all taken`, func(t *testing.T) {
		player := PlayerX
		g := NewGrid()

		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				_ = g.Take(player, Field{x, y})
			}
		}

		require.True(t, g.AllTaken())
	})

	t.Run(`all taken except one`, func(t *testing.T) {
		player := PlayerX
		g := NewGrid()

		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				if x == 2 && y == 2 {
					continue
				}
				_ = g.Take(player, Field{x, y})
			}
		}

		require.False(t, g.AllTaken())
	})
}

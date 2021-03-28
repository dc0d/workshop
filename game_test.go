package workshop

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Game_player_X_can_not_play_two_consecutive_turns(t *testing.T) {
	player := PlayerX
	field := Field{1, 1}
	game := NewGame()

	err := game.Play(player, field)
	require.NoError(t, err)

	err = game.Play(player, field)
	require.Equal(t, ErrPlayerAlreadyPlayed, err)
}

func Test_Game_a_taken_field_can_not_be_retaken(t *testing.T) {
	field := Field{1, 1}
	game := NewGame()

	err := game.Play(PlayerX, field)
	require.NoError(t, err)

	err = game.Play(PlayerO, field)
	require.Equal(t, ErrFieldTaken, err)
}

func Test_Game_a_game_is_over_when_there_is_a_winner(t *testing.T) {
	game := NewGame()

	_ = game.Play(PlayerX, Field{0, 0})
	_ = game.Play(PlayerO, Field{1, 0})
	_ = game.Play(PlayerX, Field{0, 1})
	_ = game.Play(PlayerO, Field{1, 1})
	_ = game.Play(PlayerX, Field{0, 2})

	err := game.Play(PlayerO, Field{0, 2})
	require.Equal(t, ErrGameOver, err)
	require.Equal(t, PlayerX, game.Winner())
}

func Test_Game_a_game_is_over_when_all_fields_are_taken(t *testing.T) {
	game := NewGame()

	_ = game.Play(PlayerX, Field{2, 2})
	_ = game.Play(PlayerO, Field{2, 1})
	_ = game.Play(PlayerX, Field{2, 0})
	_ = game.Play(PlayerO, Field{1, 2})
	_ = game.Play(PlayerX, Field{1, 1})
	_ = game.Play(PlayerO, Field{0, 2})
	_ = game.Play(PlayerX, Field{1, 0})
	_ = game.Play(PlayerO, Field{0, 0})
	_ = game.Play(PlayerX, Field{0, 1})

	err := game.Play(PlayerO, Field{0, 2})
	require.Equal(t, ErrGameOver, err)
	require.Equal(t, None, game.Winner())
}

package workshop_test

import (
	"testing"

	. "github.com/dc0d/workshop"

	. "github.com/dc0d/workshop/bddl"
	assert "github.com/stretchr/testify/require"
)

func Test_game_suite(t *testing.T) {
	t.Run("initial design", initialDesign)
	t.Run("partial knock down turn", partialKnockDownTurn)
	t.Run("spare turn", spareTurn)
	t.Run("strike turn", strikeTurn)
	t.Run("game total score", gameTotalScore)
	t.Run("game total score when tenth turn is a spare", gameTotalScoreWhenTenthTurnIsASpare)
	t.Run("game total score when tenth turn is a strike", gameTotalScoreWhenTenthTurnIsAStrike)
}

func initialDesign(t *testing.T) {
	var (
		game Game

		firstRoll  Roll
		secondRoll Roll

		expectedTotalScore Score
	)

	Describe(`in a game`)
	When(`game starts`, func() { game = NewGame() })
	And(`first two consecutive rolls are zero`, func() {
		firstRoll = 0
		secondRoll = 0
	})
	Doing(func() { game.PlayTurn(firstRoll, secondRoll) })
	Then(`total score should be zero`, func() {
		expectedTotalScore = 0

		assert.Equal(t, expectedTotalScore, game.Score())
	})
}

func partialKnockDownTurn(t *testing.T) {
	var (
		turn Turn

		firstRoll  Roll
		secondRoll Roll

		expectedTotalScore Score
	)

	Describe(`in each turn`)
	Given(`a new turn`, func() { turn = NewTurn() })
	When(`in two tries, the player fails to knock all pins down`, func() {
		firstRoll = 5
		secondRoll = 4
	})
	Doing(func() { turn.Play(firstRoll, secondRoll) })
	Then(`score for that turn is the total number of pins knocked down`, func() {
		expectedTotalScore = Score(firstRoll + secondRoll)

		assert.Equal(t, expectedTotalScore, turn.Score())
	})
}

func spareTurn(t *testing.T) {
	var (
		turn Turn

		firstRoll  Roll
		secondRoll Roll

		firstRollNextTurn Roll

		expectedTotalScore Score
	)

	Describe(`in a turn`, func() { turn = NewTurn() })
	When(`in two tries, the player knocks all pins down`, func() {
		firstRoll = 6
		secondRoll = 4
	})
	Doing(func() { turn.Play(firstRoll, secondRoll) })
	Doing(func() {
		nextTurn := NewTurn()

		firstRollNextTurn = 2
		nextTurn.Play(firstRollNextTurn, 0)

		turn.Next(nextTurn)
	})
	Then(`the score of the turn is ten plus the first roll of the next turn`, func() {
		expectedTotalScore = Score(firstRoll + secondRoll + firstRollNextTurn)

		assert.Equal(t, expectedTotalScore, turn.Score())
	})
}

func strikeTurn(t *testing.T) {
	var (
		turn Turn

		firstRoll  Roll
		secondRoll Roll

		firstRollNextTurn  Roll
		secondRollNextTurn Roll

		expectedTotalScore Score
	)

	Describe(`in a turn`, func() { turn = NewTurn() })
	When(`on first roll in a turn, the player knocks down all the pins`, func() {
		firstRoll = 10
		secondRoll = 0
	})
	Doing(func() { turn.Play(firstRoll, secondRoll) })
	Doing(func() {
		nextTurn := NewTurn()

		firstRollNextTurn = 3
		secondRollNextTurn = 6
		nextTurn.Play(firstRollNextTurn, secondRollNextTurn)

		turn.Next(nextTurn)
	})
	Then(`the score of the turn is ten plus the simple total of the pins knocked down in next two rolls`, func() {
		expectedTotalScore = Score(firstRoll + secondRoll + firstRollNextTurn + secondRollNextTurn)

		assert.Equal(t, expectedTotalScore, turn.Score())
	})
}

func gameTotalScore(t *testing.T) {
	type turnRolls struct{ first, second Roll }

	var (
		game  Game
		turns [10]turnRolls

		expectedTotalScore Score
	)

	Describe(`in a game`, func() { game = NewGame() })
	Given(`10 turns`, func() {
		turns[0] = turnRolls{1, 1}
		turns[1] = turnRolls{6, 4}
		turns[2] = turnRolls{1, 1}
		turns[3] = turnRolls{1, 1}
		turns[4] = turnRolls{10, 0}
		turns[5] = turnRolls{1, 1}
		turns[6] = turnRolls{1, 1}
		turns[7] = turnRolls{1, 1}
		turns[8] = turnRolls{1, 1}
		turns[9] = turnRolls{1, 1}
	})
	Doing(func() {
		for _, v := range turns {
			game.PlayTurn(v.first, v.second)
		}
	})
	Then(`the game score is the total of all turn scores`, func() {
		expectedTotalScore =
			(1 + 1) +
				(6 + 4 + 1) + // spare
				(1 + 1) +
				(1 + 1) +
				(10 + 1 + 1) + // strike
				(1 + 1) +
				(1 + 1) +
				(1 + 1) +
				(1 + 1) +
				(1 + 1)

		assert.Equal(t, expectedTotalScore, game.Score())
	})
}

func gameTotalScoreWhenTenthTurnIsASpare(t *testing.T) {
	type turnRolls struct{ first, second Roll }

	var (
		game  Game
		turns [11]turnRolls

		expectedTotalScore Score
	)

	Describe(`in a game`, func() { game = NewGame() })
	Given(`10 turns in which the last is a spare`, func() {
		turns[0] = turnRolls{1, 1}
		turns[1] = turnRolls{6, 4}
		turns[2] = turnRolls{1, 1}
		turns[3] = turnRolls{1, 1}
		turns[4] = turnRolls{10, 0}
		turns[5] = turnRolls{1, 1}
		turns[6] = turnRolls{1, 1}
		turns[7] = turnRolls{1, 1}
		turns[8] = turnRolls{1, 1}
		turns[9] = turnRolls{1, 9}
		turns[10] = turnRolls{1, 0}
	})
	Doing(func() {
		for _, v := range turns {
			game.PlayTurn(v.first, v.second)
		}
	})
	Then(`the game score is the total of all turn scores plus an extra roll`, func() {
		expectedTotalScore =
			(1 + 1) +
				(6 + 4 + 1) + // spare
				(1 + 1) +
				(1 + 1) +
				(10 + 1 + 1) + // strike
				(1 + 1) +
				(1 + 1) +
				(1 + 1) +
				(1 + 1) +
				(1 + 9 + 1) // spare at the end

		assert.Equal(t, expectedTotalScore, game.Score())
	})
}

func gameTotalScoreWhenTenthTurnIsAStrike(t *testing.T) {
	type turnRolls struct{ first, second Roll }

	var (
		game  Game
		turns [11]turnRolls

		expectedTotalScore Score
	)

	Describe(`in a game`, func() { game = NewGame() })
	Given(`10 turns in which the last is a strike`, func() {
		turns[0] = turnRolls{1, 1}
		turns[1] = turnRolls{6, 4}
		turns[2] = turnRolls{1, 1}
		turns[3] = turnRolls{1, 1}
		turns[4] = turnRolls{10, 0}
		turns[5] = turnRolls{1, 1}
		turns[6] = turnRolls{1, 1}
		turns[7] = turnRolls{1, 1}
		turns[8] = turnRolls{1, 1}
		turns[9] = turnRolls{10, 0}
		turns[10] = turnRolls{3, 6}
	})
	Doing(func() {
		for _, v := range turns {
			game.PlayTurn(v.first, v.second)
		}
	})
	Then(`the game score is the total of all turn scores plus two extra rolls`, func() {
		expectedTotalScore =
			(1 + 1) +
				(6 + 4 + 1) + // spare
				(1 + 1) +
				(1 + 1) +
				(10 + 1 + 1) + // strike
				(1 + 1) +
				(1 + 1) +
				(1 + 1) +
				(1 + 1) +
				(10 + 3 + 6) // strike at the end

		assert.Equal(t, expectedTotalScore, game.Score())
	})
}

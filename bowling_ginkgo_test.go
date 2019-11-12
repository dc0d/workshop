package workshop_test

import (
	"testing"

	. "github.com/dc0d/workshop"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func _TestGame(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Game Suite")
}

var _ = Describe("Game", func() {
	var (
		game Game

		firstRoll  Roll
		secondRoll Roll

		expectedTotalScore Score
	)

	When(`game starts`, func() {
		game = NewGame()

		Context(`and first two consecutive rolls are zero`, func() {
			firstRoll = 0
			secondRoll = 0

			It(`total score should be zero`, func() {
				expectedTotalScore = 0

				game.PlayTurn(firstRoll, secondRoll)

				Expect(game.Score()).To(Equal(expectedTotalScore))
			})
		})
	})
})

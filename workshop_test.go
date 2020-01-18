package workshop_test

import (
	"fmt"
	"testing"

	"github.com/dc0d/workshop"

	assert "github.com/stretchr/testify/require"
)

func Test_rover_turn(t *testing.T) {
	type (
		expectation struct {
			input          string
			expectedOutput string
		}
	)

	var (
		expectations = []expectation{
			{"R", "0:0:E"},
			{"RR", "0:0:S"},
			{"RRR", "0:0:W"},
			{"RRRR", "0:0:N"},
			{"L", "0:0:W"},
			{"LL", "0:0:S"},
			{"LLL", "0:0:E"},
			{"LLLL", "0:0:N"},
			{"LRLR", "0:0:N"},
		}
	)

	for _, exp := range expectations {
		var (
			commands         = exp.input
			expectedPosition = exp.expectedOutput
		)

		t.Run(fmt.Sprintf("roevr executing commands %v should go to %v", commands, expectedPosition), func(t *testing.T) {
			var (
				grid  *workshop.Grid
				rover *workshop.Rover

				assert = assert.New(t)
			)

			{
				grid = workshop.NewGrid(10, 10)
				rover = workshop.NewRover(grid)
			}

			position := rover.Execute(commands)

			assert.Equal(expectedPosition, position)
		})
	}
}

func Test_rover_move(t *testing.T) {
	type (
		expectation struct {
			input          string
			expectedOutput string
		}
	)

	var (
		expectations = []expectation{
			{"M", "0:1:N"},
			{"RM", "1:0:E"},
			{"RMM", "2:0:E"},
			{"MMRMMRMRM", "1:1:W"},
			{"MM", "0:2:N"},
			{"MMMMMMMMMM", "0:0:N"},
			{"RMMMMMMMMMM", "0:0:E"},
			{"MMMMMMMMMRRMMMMMMMMMM", "0:9:S"},
			{"RMMMMMMMMMRRMMMMMMMMMM", "9:0:W"},
			{"RRM", "0:9:S"},
			{"LM", "9:0:W"},
		}
	)

	for _, exp := range expectations {
		var (
			commands         = exp.input
			expectedPosition = exp.expectedOutput
		)

		t.Run(fmt.Sprintf("roevr executing commands %v should go to %v", commands, expectedPosition), func(t *testing.T) {
			var (
				grid  *workshop.Grid
				rover *workshop.Rover

				assert = assert.New(t)
			)

			{
				grid = workshop.NewGrid(10, 10)
				rover = workshop.NewRover(grid)
			}

			position := rover.Execute(commands)

			assert.Equal(expectedPosition, position)
		})
	}
}

func Test_rover_move_obstacles(t *testing.T) {
	var (
		grid  *workshop.Grid
		rover *workshop.Rover

		assert = assert.New(t)
	)

	{
		grid = workshop.NewGrid(10, 10)
		grid.AddObstacle(workshop.NewPosition(2, 2))

		rover = workshop.NewRover(grid)
	}

	position := rover.Execute("MMMRMMRM")
	assert.Equal("O:2:2:S", position)
}

func Test_grid(t *testing.T) {
	var (
		grid *workshop.Grid

		p4 = workshop.NewPosition(0, 10)
		p5 = workshop.NewPosition(10, 0)
	)

	{
		grid = workshop.NewGrid(10, 10)
	}

	if workshop.ErrOutside != grid.AddObstacle(p4) {
		t.FailNow()
	}

	if workshop.ErrOutside != grid.AddObstacle(p5) {
		t.FailNow()
	}
}

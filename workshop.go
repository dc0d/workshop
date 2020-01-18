package workshop

import (
	"errors"
	"fmt"
)

type Rover struct {
	grid      *Grid
	direction direction
	position  Position
}

func NewRover(grid *Grid) *Rover {
	res := &Rover{
		grid:      grid,
		direction: north,
	}

	return res
}

func (r *Rover) Execute(commands string) string {
	for _, v := range commands {
		switch v {
		case 'R':
			r.direction = r.direction.Right()
		case 'L':
			r.direction = r.direction.Left()
		case 'M':
			nextPosition, hitObstacle := r.grid.moveOn(r.position, r.direction)
			if hitObstacle {
				return fmt.Sprintf("O:%d:%d:%c", nextPosition.x, nextPosition.y, r.direction)
			}
			r.position = nextPosition
		}
	}

	return fmt.Sprintf("%d:%d:%c", r.position.x, r.position.y, r.direction)
}

type Position struct{ x, y int }

func NewPosition(x, y int) Position { return Position{x: x, y: y} }

type direction rune

func (d direction) Left() direction  { return directions[d].left }
func (d direction) Right() direction { return directions[d].right }

var (
	directions = map[direction]sides{
		north: sides{left: west, right: east},
		east:  sides{left: north, right: south},
		south: sides{left: east, right: west},
		west:  sides{left: south, right: north},
	}
)

const (
	north direction = 'N'
	east  direction = 'E'
	south direction = 'S'
	west  direction = 'W'
)

type sides struct{ left, right direction }

type Grid struct {
	width     int
	height    int
	obstacles map[Position]struct{}
}

func NewGrid(width, height int) *Grid {
	res := &Grid{
		width:     width,
		height:    height,
		obstacles: make(map[Position]struct{}),
	}

	return res
}

func (g *Grid) AddObstacle(obstacles ...Position) error {
	for _, ob := range obstacles {
		if ob.x < 0 || g.width <= ob.x {
			return ErrOutside
		}
		if ob.y < 0 || g.height <= ob.y {
			return ErrOutside
		}
		g.obstacles[ob] = struct{}{}
	}
	return nil
}

func (g *Grid) isObstacle(p Position) bool {
	_, ok := g.obstacles[p]
	return ok
}

func (g *Grid) moveOn(p Position, d direction) (nextPosition Position, hitObstacle bool) {
	switch d {
	case north:
		p.y++
		if p.y >= g.height {
			p.y = 0
		}
	case east:
		p.x++
		if p.x >= g.width {
			p.x = 0
		}
	case south:
		p.y--
		if p.y < 0 {
			p.y = g.height - 1
		}
	case west:
		p.x--
		if p.x < 0 {
			p.x = g.width - 1
		}
	}

	if g.isObstacle(p) {
		return p, true
	}

	return p, false
}

var (
	ErrOutside = errors.New("position is outside the grid")
)

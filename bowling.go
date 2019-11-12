package workshop

type Game struct {
	turns [11]*Turn
}

func NewGame() (res Game) { return }

func (g *Game) PlayTurn(first, second Roll) {
	turn := NewTurn()
	turn.Play(first, second)

	var previousTurn *Turn
	for i, turn := range g.turns {
		if turn != nil {
			previousTurn = turn
			continue
		}
		t := NewTurn()
		t.Play(first, second)
		if previousTurn != nil {
			previousTurn.Next(t)
		}
		g.turns[i] = &t
		return
	}
}

func (g Game) Score() (total Score) {
	for i, turn := range g.turns {
		if turn == nil {
			return
		}
		if i == 10 {
			return
		}
		total += turn.Score()
	}
	return
}

// Roll is a number between 0 and 10 - inclusive range - indicates number of knocked-down pins
type Roll int

type Score int

type Turn struct {
	first  Roll
	second Roll

	next *Turn
}

func NewTurn() (res Turn) { return }

func (t *Turn) Play(first, second Roll) {
	t.first = first
	t.second = second
}

func (t Turn) Score() (total Score) {
	total = Score(t.first + t.second)
	if t.next == nil {
		return
	}

	if total < 10 {
		return
	}

	total += Score(t.next.first)
	if t.first == 10 {
		total += Score(t.next.second)
	}

	return
}

func (t *Turn) Next(next Turn) { t.next = &next }

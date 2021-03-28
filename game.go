package workshop

type Game struct {
	lastPlayer Player
	winner     Player
	grid       Grid
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Play(p Player, f Field) error {
	if g.IsGameOver() {
		return ErrGameOver
	}

	if !g.canPlay(p) {
		return ErrPlayerAlreadyPlayed
	}

	if err := g.grid.Take(p, f); err != nil {
		return err
	}

	g.lastPlayer = p

	return nil
}

func (g *Game) Winner() Player { return g.winner }

func (g *Game) IsGameOver() bool {
	if g.grid.AllTaken() {
		return true
	}
	if g.winner == None {
		g.winner = g.grid.Winner()
	}
	if g.winner != None {
		return true
	}

	return false
}

func (g *Game) canPlay(p Player) bool {
	return p != g.lastPlayer
}

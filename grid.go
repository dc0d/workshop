package workshop

type Grid [3][3]Player

func NewGrid() Grid {
	var g Grid
	return g
}

func (g *Grid) Take(p Player, f Field) error {
	if !p.IsValid() {
		return ErrPlayerInvalid
	}
	if !f.IsValid() {
		return ErrFieldInvalid
	}
	if g.isTaken(f) {
		return ErrFieldTaken
	}

	g[f.X][f.Y] = p

	return nil
}

func (g Grid) AllTaken() bool {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if g[x][y] == None {
				return false
			}
		}
	}
	return true
}

func (g Grid) Winner() Player {
	if p := g.wonByColumn(); p != None {
		return p
	}

	if p := g.wonByRow(); p != None {
		return p
	}

	if p := g.wonByDiagonal(); p != None {
		return p
	}

	return None
}

func (g Grid) wonByDiagonal() Player {
	wx, wo := winners()

	var (
		d1, d2 [3]Player
	)

	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if x != y {
				continue
			}

			d1[x] = g[x][y]
			d2[x] = g[x][2-y]
		}
	}

	switch {
	case d1 == wx || d2 == wx:
		return PlayerX
	case d1 == wo || d2 == wo:
		return PlayerO
	}

	return None
}

func (g Grid) wonByRow() Player {
	wx, wo := winners()

	for y := 0; y < 3; y++ {
		var row [3]Player
		for x := 0; x < 3; x++ {
			row[x] = g[x][y]
		}
		if row == wx {
			return PlayerX
		}
		if row == wo {
			return PlayerO
		}
	}

	return None
}

func (g Grid) wonByColumn() Player {
	wx, wo := winners()

	for x := 0; x < 3; x++ {
		if g[x] == wx {
			return PlayerX
		}
		if g[x] == wo {
			return PlayerO
		}
	}

	return None
}

func (g Grid) isTaken(f Field) bool {
	return g[f.X][f.Y] != None
}

func winners() (wx, wo [3]Player) {
	for i := 0; i < 3; i++ {
		wx[i] = PlayerX
		wo[i] = PlayerO
	}
	return
}

package workshop

type Player int

func (p Player) IsValid() bool {
	if p == PlayerX || p == PlayerO {
		return true
	}

	return false
}

const (
	None Player = iota
	PlayerX
	PlayerO
)

package workshop

import (
	"math"
)

type parrotType int

const (
	TypeEuropean      parrotType = 1
	TypeAfrican       parrotType = 2
	TypeNorwegianBlue parrotType = 3
)

// Parrot has a Speed.
type Parrot interface {
	Speed() (float64, error)
}

func CreateParrot(t parrotType, numberOfCoconuts int, voltage float64, nailed bool) Parrot {
	if t == TypeEuropean {
		return typeEuropean{}
	}

	if t == TypeAfrican {
		return typeAfrican{
			numberOfCoconuts: numberOfCoconuts,
		}
	}

	if t == TypeNorwegianBlue {
		return typeNorwegianBlue{
			voltage: voltage,
			nailed:  nailed,
		}
	}

	panic("not implemented")
}

type parrot struct{}

func (parrot) baseSpeed() float64 { return 12.0 }

type typeEuropean struct {
	parrot
}

func (parrot typeEuropean) Speed() (float64, error) { return parrot.baseSpeed(), nil }

type typeAfrican struct {
	parrot

	numberOfCoconuts int
}

func (typeAfrican) loadFactor() float64 { return 9.0 }

func (parrot typeAfrican) Speed() (float64, error) {
	return math.Max(0, parrot.baseSpeed()-parrot.loadFactor()*float64(parrot.numberOfCoconuts)), nil
}

type typeNorwegianBlue struct {
	parrot

	voltage float64
	nailed  bool
}

func (parrot typeNorwegianBlue) computeBaseSpeedForVoltage(voltage float64) float64 {
	return math.Min(24.0, voltage*parrot.baseSpeed())
}

func (parrot typeNorwegianBlue) Speed() (float64, error) {
	if parrot.nailed {
		return 0, nil
	}
	return parrot.computeBaseSpeedForVoltage(parrot.voltage), nil
}

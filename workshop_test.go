package workshop_test

import (
	"testing"

	"github.com/dc0d/workshop"

	assert "github.com/stretchr/testify/require"
)

var _ *assert.Assertions

func TestSpeedOfEuropeanParrot(t *testing.T) {
	var (
		parrot workshop.Parrot

		assert = assert.New(t)
	)

	{
		parrot = workshop.CreateParrot(workshop.TypeEuropean, 0, 0, false)
	}

	speed, err := parrot.Speed()

	assert.Equal(12.0, speed)
	assert.NoError(err)
}

func TestSpeedOfAfricanParrot_With_One_Coconut(t *testing.T) {
	var (
		parrot workshop.Parrot

		assert = assert.New(t)
	)

	{
		parrot = workshop.CreateParrot(workshop.TypeAfrican, 1, 0, false)
	}

	speed, err := parrot.Speed()

	assert.Equal(3.0, speed)
	assert.NoError(err)
}

func TestSpeedOfAfricanParrot_With_Two_Coconuts(t *testing.T) {
	var (
		parrot workshop.Parrot

		assert = assert.New(t)
	)

	{
		parrot = workshop.CreateParrot(workshop.TypeAfrican, 2, 0, false)
	}

	speed, err := parrot.Speed()

	assert.Equal(0.0, speed)
	assert.NoError(err)
}

func TestSpeedOfAfricanParrot_With_No_Coconuts(t *testing.T) {
	var (
		parrot workshop.Parrot

		assert = assert.New(t)
	)

	{
		parrot = workshop.CreateParrot(workshop.TypeAfrican, 0, 0, false)
	}

	speed, err := parrot.Speed()

	assert.Equal(12.0, speed)
	assert.NoError(err)
}

func TestSpeedNorwegianBlueParrot_nailed(t *testing.T) {
	var (
		parrot workshop.Parrot

		assert = assert.New(t)
	)

	{
		parrot = workshop.CreateParrot(workshop.TypeNorwegianBlue, 0, 1.5, true)
	}

	speed, err := parrot.Speed()

	assert.Equal(0.0, speed)
	assert.NoError(err)
}

func TestSpeedNorwegianBlueParrot_not_nailed(t *testing.T) {
	var (
		parrot workshop.Parrot

		assert = assert.New(t)
	)

	{
		parrot = workshop.CreateParrot(workshop.TypeNorwegianBlue, 0, 1.5, false)
	}

	speed, err := parrot.Speed()

	assert.Equal(18.0, speed)
	assert.NoError(err)
}

func TestSpeedNorwegianBlueParrot_not_nailed_high_voltage(t *testing.T) {
	var (
		parrot workshop.Parrot

		assert = assert.New(t)
	)

	{
		parrot = workshop.CreateParrot(workshop.TypeNorwegianBlue, 0, 4, false)
	}

	speed, err := parrot.Speed()

	assert.Equal(24.0, speed)
	assert.NoError(err)
}

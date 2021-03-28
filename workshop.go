package workshop

var (
	ErrFieldTaken          sentinelError = "field is taken"
	ErrFieldInvalid        sentinelError = "field is invalid"
	ErrPlayerInvalid       sentinelError = "player is invalid"
	ErrPlayerAlreadyPlayed sentinelError = "player already played"
	ErrGameOver            sentinelError = "game over"
)

type sentinelError string

func (se sentinelError) Error() string { return string(se) }

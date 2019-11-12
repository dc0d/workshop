package model

import "time"

// TimeSource .
type TimeSource interface {
	NowUTC() time.Time
}

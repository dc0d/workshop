package model

import "time"

type TimeSource interface {
	NowUTC() time.Time
}

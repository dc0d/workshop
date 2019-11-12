package repositories

import (
	"time"

	"gitlab.com/dc0d/go-workshop/model"
)

type timeSource func() time.Time

// NowUTC .
func (src timeSource) NowUTC() time.Time { return src() }

// NewTimeSource .
func NewTimeSource() model.TimeSource {
	var src timeSource
	src = func() time.Time {
		return time.Now().UTC()
	}
	return src
}

package repositories

import (
	"time"

	"github.com/dc0d/workshop/model"
)

type timeSource func() time.Time

func (src timeSource) NowUTC() time.Time { return src() }

func NewTimeSource() model.TimeSource {
	var src timeSource
	src = func() time.Time {
		return time.Now().UTC()
	}
	return src
}

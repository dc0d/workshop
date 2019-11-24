package repositories

import (
	"time"

	model "github.com/dc0d/workshop/domain_model"
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

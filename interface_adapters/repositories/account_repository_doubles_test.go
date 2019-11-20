package repositories_test

import (
	"time"

	"github.com/dc0d/workshop/model"
)

// EventStorage spy

var _ model.EventStorage = newESSpy(
	func(events ...model.EventRecord) error { panic("N/A") },
	func(streamID string) ([]model.EventRecord, error) { panic("N/A") })

type esSpy struct {
	onAppend func(events ...model.EventRecord) error
	onLoad   func(streamID string) ([]model.EventRecord, error)
}

func newESSpy(
	onAppend func(events ...model.EventRecord) error,
	onLoad func(streamID string) ([]model.EventRecord, error)) *esSpy {
	return &esSpy{
		onAppend: onAppend,
		onLoad:   onLoad,
	}
}

func (spy *esSpy) Load(streamID string) ([]model.EventRecord, error) {
	return spy.onLoad(streamID)
}

func (spy *esSpy) Append(events ...model.EventRecord) error {
	return spy.onAppend(events...)
}

// TimeSource spy

var _ model.TimeSource = newTSSpy(func() time.Time { panic("N/A") })

type tsSpy struct {
	nowUTC func() time.Time
}

func newTSSpy(nowUTC func() time.Time) *tsSpy {
	return &tsSpy{nowUTC: nowUTC}
}

func (ts *tsSpy) NowUTC() time.Time {
	return ts.nowUTC()
}

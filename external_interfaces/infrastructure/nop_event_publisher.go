package infrastructure

import "github.com/dc0d/workshop/model"

// NopEventPublisher .
type NopEventPublisher struct{}

// NewNopEventPublisher .
func NewNopEventPublisher() *NopEventPublisher {
	return &NopEventPublisher{}
}

// Publish .
func (*NopEventPublisher) Publish(...model.EventRecord) error { return nil }

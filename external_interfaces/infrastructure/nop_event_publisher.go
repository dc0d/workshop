package infrastructure

import "gitlab.com/dc0d/go-workshop/model"

// NopEventPublisher .
type NopEventPublisher struct{}

// NewNopEventPublisher .
func NewNopEventPublisher() *NopEventPublisher {
	return &NopEventPublisher{}
}

// Publish .
func (*NopEventPublisher) Publish(...model.EventRecord) error { return nil }

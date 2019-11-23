package infrastructure

import "github.com/dc0d/workshop/model"

type NopEventPublisher struct{}

func NewNopEventPublisher() *NopEventPublisher {
	return &NopEventPublisher{}
}

func (*NopEventPublisher) Publish(...model.EventRecord) error { return nil }

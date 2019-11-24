package infrastructure

import model "github.com/dc0d/workshop/domain_model"

type NopEventPublisher struct{}

func NewNopEventPublisher() *NopEventPublisher {
	return &NopEventPublisher{}
}

func (*NopEventPublisher) Publish(...model.EventRecord) error { return nil }

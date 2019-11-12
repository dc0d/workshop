package infrastructure_test

//go:generate moq -pkg infrastructure_test -out ./mock_event_queue_test.go ./../../model EventQueue

import (
	"testing"

	"gitlab.com/dc0d/go-workshop/external_interfaces/infrastructure"
	"gitlab.com/dc0d/go-workshop/model"

	"github.com/stretchr/testify/require"
)

func Test_create_queue_event_publisher(t *testing.T) {
	var queue model.EventQueue
	var _ model.EventPublisher = infrastructure.NewQueueEventPublisher(queue)
}

func Test_queue_event_publisher_publish(t *testing.T) {
	var (
		assert = require.New(t)
		id     = "ID"
		events = []model.EventRecord{
			model.EventRecord{
				StreamID: id,
				Version:  0,
			},
			model.EventRecord{
				StreamID: id,
				Version:  1,
			},
			model.EventRecord{
				StreamID: id,
				Version:  2,
			},
		}
		queuePublishCalled = false
	)

	queue := &EventQueueMock{
		PublishFunc: func(received ...model.EventRecord) error {
			queuePublishCalled = true
			assert.EqualValues(events, received)
			return nil
		},
	}
	publisher := infrastructure.NewQueueEventPublisher(queue)

	publisher.Publish(events...)

	assert.True(queuePublishCalled)
}

package infrastructure_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/dc0d/workshop/external_interfaces/infrastructure"
	model "github.com/dc0d/workshop/domain_model"
)

func Test_create_fake_event_queue(t *testing.T) {
	var _ model.EventQueue = infrastructure.NewFakeEventQueue()
}

func Test_fake_event_queue_publish_consume(t *testing.T) {
	var (
		queue  model.EventQueue
		assert = require.New(t)
	)

	queue = infrastructure.NewFakeEventQueue()

	var event model.EventRecord
	queue.Publish(event)

	messages := queue.Consume()

	select {
	case msg := <-messages:
		assert.EqualValues(event, msg)
	case <-time.After(time.Second * 3):
		t.Fatal("no message received")
	}
}

func Test_fake_event_queue_publish_consume_multiple_events(t *testing.T) {
	var (
		queue  model.EventQueue
		assert = require.New(t)
		id     = "ID"
	)

	queue = infrastructure.NewFakeEventQueue()

	events := []model.EventRecord{
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

	queue.Publish(events...)

	messages := queue.Consume()

	select {
	case msg := <-messages:
		assert.EqualValues(events[0], msg)
	case <-time.After(time.Second * 3):
		t.Fatal("no message received")
	}

	select {
	case msg := <-messages:
		assert.EqualValues(events[1], msg)
	case <-time.After(time.Second * 3):
		t.Fatal("no message received")
	}

	select {
	case msg := <-messages:
		assert.EqualValues(events[2], msg)
	case <-time.After(time.Second * 3):
		t.Fatal("no message received")
	}
}

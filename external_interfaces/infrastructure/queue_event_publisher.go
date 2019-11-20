package infrastructure

import "github.com/dc0d/workshop/model"

// QueueEventPublisher .
type QueueEventPublisher struct {
	queue model.EventQueue
}

// NewQueueEventPublisher .
func NewQueueEventPublisher(queue model.EventQueue) *QueueEventPublisher {
	return &QueueEventPublisher{
		queue: queue,
	}
}

// Publish .
func (publisher *QueueEventPublisher) Publish(events ...model.EventRecord) error {
	return publisher.queue.Publish(events...)
}

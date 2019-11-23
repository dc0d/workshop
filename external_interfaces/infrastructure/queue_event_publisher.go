package infrastructure

import "github.com/dc0d/workshop/model"

type QueueEventPublisher struct {
	queue model.EventQueue
}

func NewQueueEventPublisher(queue model.EventQueue) *QueueEventPublisher {
	return &QueueEventPublisher{
		queue: queue,
	}
}

func (publisher *QueueEventPublisher) Publish(events ...model.EventRecord) error {
	return publisher.queue.Publish(events...)
}

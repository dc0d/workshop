package infrastructure

import model "github.com/dc0d/workshop/domain_model"

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

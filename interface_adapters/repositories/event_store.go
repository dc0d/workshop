package repositories

import (
	"github.com/dc0d/workshop/model"
)

type eventStore struct {
	storage   model.EventStorage
	publisher model.EventPublisher
}

// NewEventStore .
func NewEventStore(storage model.EventStorage, publisher model.EventPublisher) model.EventStorage {
	return &eventStore{
		storage:   storage,
		publisher: publisher,
	}
}

func (store *eventStore) Load(streamID string) ([]model.EventRecord, error) {
	return store.storage.Load(streamID)
}

func (store *eventStore) Append(events ...model.EventRecord) error {
	err := store.storage.Append(events...)
	if err == nil {
		store.publisher.Publish(events...)
	}
	return err
}

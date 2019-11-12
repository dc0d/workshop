package infrastructure

import (
	"sort"
	"sync"

	"gitlab.com/dc0d/go-workshop/model"
)

type eventStorage struct {
	mx      sync.RWMutex
	storage map[string]map[int]string
}

// NewEventStorage .
func NewEventStorage() model.EventStorage {
	return &eventStorage{storage: make(map[string]map[int]string)}
}

func (storage *eventStorage) Load(streamID string) ([]model.EventRecord, error) {
	storage.mx.RLock()
	defer storage.mx.RUnlock()

	stream, ok := storage.storage[streamID]
	if !ok {
		return nil, model.ErrStreamNotFound
	}

	var res []model.EventRecord
	for version, data := range stream {
		event := model.EventRecord{
			StreamID: streamID,
			Version:  version,
			Data:     []byte(data),
		}
		res = append(res, event)
	}

	var events sortableEvents = res
	sort.Sort(events)
	res = events

	return res, nil
}

func (storage *eventStorage) Append(events ...model.EventRecord) error {
	storage.mx.Lock()
	defer storage.mx.Unlock()

	if err := storage.checkDuplicateVersion(events...); err != nil {
		return err
	}

	for _, e := range events {
		stream, ok := storage.storage[e.StreamID]
		if !ok {
			stream = make(map[int]string)
		}
		stream[e.Version] = string(e.Data)
		storage.storage[e.StreamID] = stream
	}

	return nil
}

func (storage *eventStorage) checkDuplicateVersion(events ...model.EventRecord) error {
	for _, e := range events {
		stream, ok := storage.storage[e.StreamID]
		if !ok {
			stream = make(map[int]string)
		}
		_, ok = stream[e.Version]
		if ok {
			return model.ErrDuplicateEventVersion
		}
	}

	return nil
}

type sortableEvents []model.EventRecord

func (events sortableEvents) Len() int           { return len(events) }
func (events sortableEvents) Less(i, j int) bool { return events[i].Version < events[j].Version }
func (events sortableEvents) Swap(i, j int)      { events[i], events[j] = events[j], events[i] }

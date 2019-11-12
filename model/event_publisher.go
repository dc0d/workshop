package model

// EventPublisher .
type EventPublisher interface {
	Publish(...EventRecord) error
}

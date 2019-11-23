package model

type EventPublisher interface {
	Publish(...EventRecord) error
}

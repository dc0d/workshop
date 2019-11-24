package model

type EventQueue interface {
	EventPublisher
	Consume() <-chan EventRecord
}

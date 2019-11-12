package model

// EventQueue .
type EventQueue interface {
	EventPublisher
	Consume() <-chan EventRecord
}

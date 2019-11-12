package infrastructure

import (
	"gitlab.com/dc0d/go-workshop/model"
)

// StatementViewBuilder .
type StatementViewBuilder struct {
	queue   model.EventQueue
	storage model.StatementViewStorage
}

// NewStatementViewBuilder .
func NewStatementViewBuilder(
	queue model.EventQueue,
	storage model.StatementViewStorage) *StatementViewBuilder {
	builder := &StatementViewBuilder{
		queue:   queue,
		storage: storage,
	}
	go builder.loop()
	return builder
}

func (builder *StatementViewBuilder) loop() {
	messages := builder.queue.Consume()
	for {
		msg := <-messages
		builder.storage.Save(msg)
	}
}

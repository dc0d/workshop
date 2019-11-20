package infrastructure

import (
	"github.com/dc0d/workshop/model"
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

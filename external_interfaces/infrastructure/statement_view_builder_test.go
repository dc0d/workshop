package infrastructure_test

//go:generate moq -pkg infrastructure_test -out ./mock_statement_view_storage_test.go ./../../model StatementViewStorage

import (
	"testing"
	"time"

	"gitlab.com/dc0d/go-workshop/external_interfaces/infrastructure"
	"gitlab.com/dc0d/go-workshop/model"

	"github.com/stretchr/testify/require"
)

func Test_create_statement_view_builder(t *testing.T) {
	t.SkipNow()

	var q model.EventQueue
	var s model.StatementViewStorage

	infrastructure.NewStatementViewBuilder(q, s)
}

func Test_statement_view_builder_loop(t *testing.T) {
	var (
		assert  = require.New(t)
		queue   *EventQueueMock
		storage *StatementViewStorageMock
	)

	ch := make(chan model.EventRecord, 10)
	queue = &EventQueueMock{
		ConsumeFunc: func() <-chan model.EventRecord { return ch },
	}
	storage = &StatementViewStorageMock{
		SaveFunc: func(events ...model.EventRecord) error { return nil },
	}

	viewBuilder := infrastructure.NewStatementViewBuilder(queue, storage)
	_ = viewBuilder

	var sampleEvent model.EventRecord
	sampleEvent.StreamID = "ID"
	ch <- sampleEvent

	time.Sleep(time.Millisecond * 100)
	for i := 0; i < 10; i++ {
		saveCalls := storage.SaveCalls()
		if len(saveCalls) > 0 {
			break
		}
		time.Sleep(time.Millisecond * 100 * time.Duration(i+1))
	}

	consumeCalls := queue.ConsumeCalls()
	assert.True(len(consumeCalls) > 0)

	saveCalls := storage.SaveCalls()
	assert.True(len(saveCalls) > 0)
	assert.EqualValues(sampleEvent, saveCalls[0].In1[0])
}

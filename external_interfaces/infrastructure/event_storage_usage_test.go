package infrastructure_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/dc0d/workshop/external_interfaces/infrastructure"
	"github.com/dc0d/workshop/model"

	"github.com/stretchr/testify/require"
)

func Test_create_event_storage(t *testing.T) {
	var _ model.EventStorage = infrastructure.NewEventStorage()
}

func Test_append(t *testing.T) {
	assert := require.New(t)

	storage := infrastructure.NewEventStorage()
	err := storage.Append(sampleEvents1()...)
	assert.NoError(err)
}

func Test_stream_not_found(t *testing.T) {
	assert := require.New(t)

	storage := infrastructure.NewEventStorage()
	events, err := storage.Load("NON_EXISTING_STREAM")

	assert.Equal(model.ErrStreamNotFound, err)
	assert.Len(events, 0)
}

func Test_append_duplicated_version(t *testing.T) {
	t.Run("append duplicated version", func(t *testing.T) {
		assert := require.New(t)

		storage := infrastructure.NewEventStorage()
		storage.Append(sampleEvents1()...)

		err := storage.Append(sampleEvents1()...)
		assert.Equal(model.ErrDuplicateEventVersion, err)
	})

	t.Run("append duplicated version and load", func(t *testing.T) {
		assert := require.New(t)

		numberOfEvents := 200

		storage := infrastructure.NewEventStorage()
		sampleEvents2 := sampleEvents2(numberOfEvents)
		storage.Append(sampleEvents2...)
		storage.Append(sampleEvents2...)
		storage.Append(sampleEvents2...)

		found, err := storage.Load("AN_ID")
		assert.NoError(err)
		assert.Equal(numberOfEvents, len(found))

		for i := 0; i < numberOfEvents; i++ {
			assert.Equal(sampleEvents2[i].StreamID, found[i].StreamID)
			assert.Equal(sampleEvents2[i].Version, found[i].Version)
			assert.Equal(sampleEvents2[i].Data, found[i].Data)
		}
	})
}

func Test_load(t *testing.T) {
	assert := require.New(t)

	numberOfEvents := 20

	storage := infrastructure.NewEventStorage()
	sampleEvents2 := sampleEvents2(numberOfEvents)
	storage.Append(sampleEvents2...)

	found, err := storage.Load("AN_ID")
	assert.NoError(err)
	assert.Equal(numberOfEvents, len(found))

	for i := 0; i < numberOfEvents; i++ {
		assert.Equal(sampleEvents2[i].StreamID, found[i].StreamID)
		assert.Equal(sampleEvents2[i].Version, found[i].Version)
		assert.Equal(sampleEvents2[i].Data, found[i].Data)
	}
}

func sampleEvents2(numberOfEvents int) (res []model.EventRecord) {
	id := "AN_ID"

	for i := 0; i < numberOfEvents; i++ {
		var domainEvent1 EventA
		domainEvent1.ID = id
		domainEvent1.Data1 = (i + 1) * 10
		domainEvent1.Data2 = fmt.Sprintf("DATA-%v", domainEvent1.Data1)
		js1 := toJSON(domainEvent1)

		e1 := model.EventRecord{
			StreamID: id,
			Version:  i,
			Data:     js1,
		}

		res = append(res, e1)
	}

	return
}

func sampleEvents1() []model.EventRecord {
	id := "AN_ID"

	var domainEvent1 EventA
	domainEvent1.ID = id
	domainEvent1.Data1 = 110
	domainEvent1.Data2 = "DATA"
	js1 := toJSON(domainEvent1)

	var domainEvent2 EventB
	domainEvent2.ID = id
	domainEvent2.Data3 = "DATA"
	js2 := toJSON(domainEvent2)

	e1 := model.EventRecord{
		StreamID: id,
		Version:  1,
		Data:     js1,
	}
	e2 := model.EventRecord{
		StreamID: id,
		Version:  2,
		Data:     js2,
	}

	return []model.EventRecord{e1, e2}
}

func toJSON(v interface{}) []byte {
	js, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return js
}

type EventA struct {
	ID    string `json:"id"`
	Data1 int    `json:"data_1,string,omitempty"`
	Data2 string `json:"data_2,omitempty"`
}

type EventB struct {
	ID    string `json:"id"`
	Data3 string `json:"data_3,omitempty"`
}

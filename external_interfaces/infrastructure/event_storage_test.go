package infrastructure

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/dc0d/workshop/model"
)

func Test_sortable_events(t *testing.T) {
	assert := require.New(t)

	var events sortableEvents
	var _ sort.Interface = events

	numberOfEvents := 100

	sampleEvents2 := sampleEvents2(numberOfEvents)
	events = make([]model.EventRecord, numberOfEvents)
	for k := range sampleEvents2 {
		events[numberOfEvents-k-1] = sampleEvents2[k]
	}

	assert.Greater(events[0].Version, events[1].Version)
	sort.Sort(events)
	for i := 0; i < numberOfEvents-1; i++ {
		assert.Less(events[i].Version, events[i+1].Version)
	}
}

func sampleEvents2(numberOfEvents int) (res []model.EventRecord) {
	id := "AN_ID"

	for i := 0; i < numberOfEvents; i++ {
		var domainEvent1 EventA
		domainEvent1.ID = id
		domainEvent1.Data1 = (i + 1) * 10
		domainEvent1.Data2 = fmt.Sprintf("DATA-%v", domainEvent1.Data1)
		js1 := toJS(domainEvent1)

		e1 := model.EventRecord{
			StreamID: id,
			Version:  i,
			Data:     js1,
		}

		res = append(res, e1)
	}

	return
}

func toJS(v interface{}) []byte {
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

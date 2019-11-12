package infrastructure

import "gitlab.com/dc0d/go-workshop/model"

// FakeEventQueue .
type FakeEventQueue struct {
	incoming chan model.EventRecord
	outgoing chan model.EventRecord
}

// NewFakeEventQueue .
func NewFakeEventQueue() *FakeEventQueue {
	queue := &FakeEventQueue{
		incoming: make(chan model.EventRecord),
		outgoing: make(chan model.EventRecord),
	}
	go queue.loop()
	return queue
}

// Publish .
func (queue *FakeEventQueue) Publish(messages ...model.EventRecord) error {
	for _, msg := range messages {
		queue.incoming <- msg
	}
	return nil
}

// Consume .
func (queue *FakeEventQueue) Consume() <-chan model.EventRecord {
	return queue.outgoing
}

func (queue *FakeEventQueue) loop() {
	var q []model.EventRecord
	for {
		var outgoing chan model.EventRecord
		var currentMessage model.EventRecord

		if len(q) > 0 {
			outgoing = queue.outgoing
			currentMessage = q[0]
		} else {
			outgoing = nil
		}

		select {
		case msg := <-queue.incoming:
			q = append(q, msg)
		case outgoing <- currentMessage:
			q = q[1:]
		}
	}
}

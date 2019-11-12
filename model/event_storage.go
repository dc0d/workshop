package model

// EventStorage .
type EventStorage interface {
	Load(streamID string) ([]EventRecord, error)
	Append(events ...EventRecord) error
}

// EventRecord .
type EventRecord struct {
	StreamID string
	Version  int
	Data     []byte
}

// EventRecordData is the actual data that is put in EventRecord.Data
type EventRecordData struct {
	Type      string
	EventData []byte
}

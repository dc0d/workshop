package model

import "time"

type StreamEvent interface {
	GetID() string
	GetTimestamp() time.Time
	SetTimestamp(time.Time)
	GetVersion() int
	SetVersion(int)
}

type streamEvent struct {
	ID        string
	Timestamp time.Time
	Version   int
}

func (e streamEvent) GetID() string             { return e.ID }
func (e streamEvent) GetTimestamp() time.Time   { return e.Timestamp }
func (e *streamEvent) SetTimestamp(t time.Time) { e.Timestamp = t }
func (e streamEvent) GetVersion() int           { return e.Version }
func (e *streamEvent) SetVersion(v int)         { e.Version = v }

type AccountCreated struct {
	streamEvent

	ClientID string
}

func (e AccountCreated) Validate() error {
	if e.GetID() == "" {
		return ErrAccountIDEmpty
	}
	if e.ClientID == "" {
		return ErrClientIDEmpty
	}
	return nil
}

type AmountDeposited struct {
	streamEvent

	Amount          Amount
	TransactionTime time.Time
}

type AmountWithdrawn struct {
	streamEvent

	Amount          Amount
	TransactionTime time.Time
}

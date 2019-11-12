package model

// StatementViewStorage .
type StatementViewStorage interface {
	Find(id string) (*Statement, error)
	Save(...EventRecord) error
}

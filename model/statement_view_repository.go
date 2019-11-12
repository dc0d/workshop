package model

// StatementViewRepository .
type StatementViewRepository interface {
	Find(id string) (*Statement, error)
}

package model

type StatementViewRepository interface {
	Find(id string) (*Statement, error)
}

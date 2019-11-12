package repositories

import "gitlab.com/dc0d/go-workshop/model"

// StatementViewRepository .
type StatementViewRepository struct {
	view StatementViewStorageRead
}

// NewStatementViewRepository .
func NewStatementViewRepository(view StatementViewStorageRead) *StatementViewRepository {
	return &StatementViewRepository{
		view: view,
	}
}

// Find .
func (repo *StatementViewRepository) Find(id string) (*model.Statement, error) {
	return repo.view.Find(id)
}

// StatementViewStorageRead .
type StatementViewStorageRead interface {
	Find(id string) (*model.Statement, error)
}

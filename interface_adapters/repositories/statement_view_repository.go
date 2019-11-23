package repositories

import "github.com/dc0d/workshop/model"

type StatementViewRepository struct {
	view StatementViewStorageRead
}

func NewStatementViewRepository(view StatementViewStorageRead) *StatementViewRepository {
	return &StatementViewRepository{
		view: view,
	}
}

func (repo *StatementViewRepository) Find(id string) (*model.Statement, error) {
	return repo.view.Find(id)
}

type StatementViewStorageRead interface {
	Find(id string) (*model.Statement, error)
}

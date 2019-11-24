package usecases

import (
	model "github.com/dc0d/workshop/domain_model"
)

type bankStatement struct {
	repo model.StatementViewRepository
}

func NewBankStatement(repo model.StatementViewRepository) model.BankStatement {
	return &bankStatement{repo: repo}
}

func (usecase *bankStatement) Run(id string) (*model.Statement, error) {
	statement, err := usecase.repo.Find(id)
	if err != nil {
		return nil, err
	}
	return statement, nil
}

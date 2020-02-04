package usecases

import (
	model "github.com/dc0d/workshop/domain_model"
)

type BankStatement struct {
	repo model.StatementViewRepository
}

func NewBankStatement(repo model.StatementViewRepository) *BankStatement {
	return &BankStatement{repo: repo}
}

func (usecase *BankStatement) Run(id string) (*model.Statement, error) {
	statement, err := usecase.repo.Find(id)
	if err != nil {
		return nil, err
	}
	return statement, nil
}

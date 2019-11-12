package usecases

import "gitlab.com/dc0d/go-workshop/model"

type bankStatement struct {
	repo model.StatementViewRepository
}

// NewBankStatement .
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

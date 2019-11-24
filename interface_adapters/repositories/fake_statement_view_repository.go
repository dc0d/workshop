package repositories

import (
	model "github.com/dc0d/workshop/domain_model"
)

type FakeStatementViewRepository struct {
	accountRepo model.AccountRepository
}

func NewFakeStatementViewRepository(accountRepo model.AccountRepository) *FakeStatementViewRepository {
	return &FakeStatementViewRepository{
		accountRepo: accountRepo,
	}
}

func (repo *FakeStatementViewRepository) Find(id string) (*model.Statement, error) {
	account, err := repo.accountRepo.Find(id)
	if err != nil {
		return nil, err
	}
	return account.Statement(), nil
}

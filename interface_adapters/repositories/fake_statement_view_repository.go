package repositories

import (
	"gitlab.com/dc0d/go-workshop/model"
)

// FakeStatementViewRepository .
type FakeStatementViewRepository struct {
	accountRepo model.AccountRepository
}

// NewFakeStatementViewRepository .
func NewFakeStatementViewRepository(accountRepo model.AccountRepository) *FakeStatementViewRepository {
	return &FakeStatementViewRepository{
		accountRepo: accountRepo,
	}
}

// Find .
func (repo *FakeStatementViewRepository) Find(id string) (*model.Statement, error) {
	account, err := repo.accountRepo.Find(id)
	if err != nil {
		return nil, err
	}
	return account.Statement(), nil
}

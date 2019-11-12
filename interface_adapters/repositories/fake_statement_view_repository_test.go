package repositories_test

//go:generate moq -pkg repositories_test -out ./mock_account_repository_test.go ./../../model AccountRepository

import (
	"testing"

	"gitlab.com/dc0d/go-workshop/interface_adapters/repositories"
	"gitlab.com/dc0d/go-workshop/model"

	"github.com/stretchr/testify/require"
)

func Test_fake_statement_view_repository_interface(t *testing.T) {
	var accountRepo model.AccountRepository
	var _ model.StatementViewRepository = repositories.NewFakeStatementViewRepository(accountRepo)
}

func Test_find_statement(t *testing.T) {
	assert := require.New(t)

	id := "ID"

	accountRepo := &AccountRepositoryMock{
		FindFunc: func(_id string) (*model.Account, error) {
			account := model.NewAccount(id)
			account.CreateAccount(id)
			account.Deposit(1000, parseDate("10-01-2012"))
			account.Deposit(2000, parseDate("13-01-2012"))
			account.Withdraw(500, parseDate("14-01-2012"))
			return account, nil
		},
	}

	repo := repositories.NewFakeStatementViewRepository(accountRepo)

	statement, err := repo.Find(id)
	assert.NoError(err)
	assert.Equal(expectedBankStatement, statement.String())

	findCalls := accountRepo.FindCalls()
	assert.Len(findCalls, 1)
}

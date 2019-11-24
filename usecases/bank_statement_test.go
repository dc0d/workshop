package usecases_test

//go:generate moq -pkg usecases_test -out ./mock_statement_view_repository_test.go ./../domain_model StatementViewRepository

import (
	"testing"

	model "github.com/dc0d/workshop/domain_model"
	"github.com/dc0d/workshop/usecases"

	"github.com/stretchr/testify/require"
)

func Test_create_bank_statement_usecase(t *testing.T) {
	var statementRepo model.StatementViewRepository
	var _ model.BankStatement = usecases.NewBankStatement(statementRepo)
}

func Test_get_bank_statement(t *testing.T) {
	assert := require.New(t)

	id := "A_CLIENT_ID"

	repo := &StatementViewRepositoryMock{}

	{
		account := model.NewAccount("")
		account.RebuildFrom(sampleEventsForBuild(id)...)
		expectedStatement := account.Statement()

		repo.FindFunc = func(id string) (*model.Statement, error) {
			return expectedStatement, nil
		}
	}

	usecase := usecases.NewBankStatement(repo)

	statement, err := usecase.Run(id)
	assert.NoError(err)
	assert.Equal(expectedBankStatement, statement.String())
}

func sampleEventsForBuild(id string) (res []model.StreamEvent) {
	{
		e := model.AccountCreated{
			ClientID: id,
		}
		e.ID = id
		e.Timestamp = parseDate("01-01-2019")
		e.Version = 0
		res = append(res, &e)
	}
	{
		e := model.AmountDeposited{
			Amount: 1000,
		}
		e.ID = id
		e.Timestamp = parseDate("10-01-2019")
		e.TransactionTime = parseDate("10-01-2012")
		e.Version = 1
		res = append(res, &e)
	}
	{
		e := model.AmountDeposited{
			Amount: 2000,
		}
		e.ID = id
		e.Timestamp = parseDate("13-01-2019")
		e.TransactionTime = parseDate("13-01-2012")
		e.Version = 2
		res = append(res, &e)
	}
	{
		e := model.AmountWithdrawn{
			Amount: 500,
		}
		e.ID = id
		e.Timestamp = parseDate("14-01-2019")
		e.TransactionTime = parseDate("14-01-2012")
		e.Version = 3
		res = append(res, &e)
	}
	return
}

var (
	expectedBankStatement = `date || credit || debit || balance
14/01/2012 || || 500.00 || 2500.00
13/01/2012 || 2000.00 || || 3000.00
10/01/2012 || 1000.00 || || 1000.00`
)

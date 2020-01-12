package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	model "github.com/dc0d/workshop/domain_model"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/require"
)

func Test_transaction_command_handler_using_the_router(t *testing.T) {
	var (
		clientID          = "A_CLIENT_ID"
		transactionDate   = parseDate("02-11-2019")
		transactionAmount = 1000

		command        transactionCommand
		commandPayload []byte
		opt            model.HandleTransactionOptions

		ctxFac ProviderContextFactory

		req    *http.Request
		rec    *httptest.ResponseRecorder
		router *echo.Echo

		usecase *HandleTransactionMock
		assert  = require.New(t)
	)

	{
		command.Command = depositCommandName
		command.Data.ClientID = clientID
		command.Data.Amount = transactionAmount
		command.Data.Time = transactionDate
		commandPayload = toJSON(command)

		req = httptest.NewRequest(
			http.MethodPost,
			"/api/bank/transactions",
			bytes.NewBuffer(commandPayload))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec = httptest.NewRecorder()

		ctxFac = func(c echo.Context) ProviderContext {
			var ctx ProviderContext

			usecase = &HandleTransactionMock{
				RunFunc: func(option model.HandleTransactionOption) error {
					option(&opt)

					return nil
				},
			}

			txFac := &HandleTransactionFactoryMock{
				CreateHandleTransactionFunc: func(model.AccountRepository) model.HandleTransaction { return usecase },
			}

			accRepoFac := &AccountRepositoryFactoryMock{
				CreateAccountRepositoryFunc: func() model.AccountRepository { return nil },
			}

			ctx = newMoqProviderContext(c, txFac, accRepoFac, nil, nil)

			return ctx
		}

		router = newRouter(ctxFac)
	}

	router.ServeHTTP(rec, req)

	assert.Empty(rec.Body.String())
	assert.Equal(http.StatusOK, rec.Code)
	assert.Condition(func() bool { return len(usecase.RunCalls()) == 1 }, "handle transaction usecase expected to be called once")
	assert.Equal(clientID, opt.DepositCommand.ClientID)
	assert.Equal(transactionAmount, opt.DepositCommand.Amount)
	assert.Equal(transactionDate, opt.DepositCommand.Time)
}

func Test_bank_statement_using_the_router(t *testing.T) {
	var (
		clientID = "A_CLIENT_ID"

		ctxFac ProviderContextFactory

		req    *http.Request
		rec    *httptest.ResponseRecorder
		router *echo.Echo

		usecase *BankStatementMock
		assert  = require.New(t)
	)

	{
		req = httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/api/bank/%v/statement", clientID),
			nil)
		rec = httptest.NewRecorder()

		ctxFac = func(c echo.Context) ProviderContext {
			var ctx ProviderContext

			usecase = &BankStatementMock{
				RunFunc: func(id string) (*model.Statement, error) {
					return sampleStatement(), nil
				},
			}

			sttRepoFac := &StatementViewRepositoryFactoryMock{
				CreateStatementViewRepositoryFunc: func() model.StatementViewRepository { return nil },
			}

			sttFac := &BankStatementFactoryMock{
				CreateBankStatementFunc: func(repo model.StatementViewRepository) model.BankStatement { return usecase },
			}

			ctx = newMoqProviderContext(c, nil, nil, sttFac, sttRepoFac)

			return ctx
		}

		router = newRouter(ctxFac)
	}

	router.ServeHTTP(rec, req)

	assert.Equal(http.StatusOK, rec.Code)
	assert.Equal(expectedBankStatement, rec.Body.String())
	assert.Condition(func() bool { return len(usecase.RunCalls()) == 1 }, "bank statement usecase expected to be called once")
	assert.Equal(clientID, usecase.RunCalls()[0].ID)
}

type moqProviderContext struct {
	model.AccountRepositoryFactory
	model.StatementViewRepositoryFactory
	model.HandleTransactionFactory
	model.BankStatementFactory

	echo.Context
}

func newMoqProviderContext(
	c echo.Context,
	txFac model.HandleTransactionFactory,
	accRepoFac model.AccountRepositoryFactory,
	sttFac model.BankStatementFactory,
	sttRepoFac model.StatementViewRepositoryFactory) ProviderContext {
	res := &moqProviderContext{
		Context:                        c,
		HandleTransactionFactory:       txFac,
		AccountRepositoryFactory:       accRepoFac,
		BankStatementFactory:           sttFac,
		StatementViewRepositoryFactory: sttRepoFac,
	}

	return res
}

func sampleStatement() *model.Statement {
	statement := model.NewStatement()

	sampleLines := sampleStatementLines()
	for _, line := range sampleLines {
		statement.AddStatementLine(line)
	}

	return statement
}

func sampleStatementLines() []model.StatementLine {
	return []model.StatementLine{
		{Date: parseDate("10-01-2012"), Credit: 1000, Debit: 0, Balance: 1000},
		{Date: parseDate("13-01-2012"), Credit: 2000, Debit: 0, Balance: 3000},
		{Date: parseDate("14-01-2012"), Credit: 0, Debit: 500, Balance: 2500},
	}
}

var (
	expectedBankStatement = `date || credit || debit || balance
14/01/2012 || || 500.00 || 2500.00
13/01/2012 || 2000.00 || || 3000.00
10/01/2012 || 1000.00 || || 1000.00`
)

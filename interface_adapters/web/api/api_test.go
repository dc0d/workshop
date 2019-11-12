package api

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"gitlab.com/dc0d/go-workshop/model"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/require"
)

func Test_transaction_command_handler_using_the_router(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	assert := require.New(t)

	clientID := "A_CLIENT_ID"
	transactionDate := parseDate("02-11-2019")
	transactionAmount := 1000

	var command transactionCommand
	command.Command = depositCommandName
	command.Data.ClientID = clientID
	command.Data.Amount = transactionAmount
	command.Data.Time = transactionDate
	commandPayload := toJSON(command)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/bank/transactions",
		bytes.NewBuffer(commandPayload))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	usecase := NewMockHandleTransaction(ctrl)
	usecase.
		EXPECT().
		Run(gomock.Any()).
		Return(nil)

	factory := newMockHandlerFactory(usecase, nil)

	router := newRouter(factory)
	router.ServeHTTP(rec, req)

	assert.Equal(http.StatusOK, rec.Code)
	assert.Empty(rec.Body.String())
}

func Test_bank_statement_using_the_router(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	assert := require.New(t)

	clientID := "A_CLIENT_ID"
	req := httptest.NewRequest(
		http.MethodGet,
		fmt.Sprintf("/api/bank/%v/statement", clientID),
		nil)
	rec := httptest.NewRecorder()

	usecase := NewMockBankStatement(ctrl)
	usecase.
		EXPECT().
		Run(clientID).
		Return(sampleStatement(), nil)

	factory := newMockHandlerFactory(nil, usecase)

	router := newRouter(factory)
	router.ServeHTTP(rec, req)

	assert.Equal(http.StatusOK, rec.Code)
	assert.Equal(expectedBankStatement, rec.Body.String())
}

type mockHandlerFactory struct {
	handleTransactionUsecase model.HandleTransaction
	bankStatementUsecase     model.BankStatement
}

func newMockHandlerFactory(
	handleTransactionUsecase model.HandleTransaction,
	bankStatementUsecase model.BankStatement) *mockHandlerFactory {
	return &mockHandlerFactory{
		handleTransactionUsecase: handleTransactionUsecase,
		bankStatementUsecase:     bankStatementUsecase,
	}
}

func (fac *mockHandlerFactory) createTransactionCommandHandler() *transactionCommandHandler {
	return newTransactionCommandHandler(fac.handleTransactionUsecase)
}

func (fac *mockHandlerFactory) createStatementHandler() *statementHandler {
	return newStatementHandler(fac.bankStatementUsecase)
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

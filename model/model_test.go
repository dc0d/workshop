package model_test

import (
	"testing"
	"time"

	"gitlab.com/dc0d/go-workshop/model"

	"github.com/stretchr/testify/require"
)

func Test_amount(t *testing.T) {
	t.Run("amount to string", func(t *testing.T) {
		assert := require.New(t)

		var a model.Amount

		a = 500
		assert.Equal("500.00", a.String())

		a = 0
		assert.Equal("", a.String())
	})
}

func Test_events(t *testing.T) {
	t.Run("account created", func(t *testing.T) {
		assert := require.New(t)

		var e model.AccountCreated

		timestamp := time.Now()
		version := 10000

		e.ID = "STREAM_ID"
		e.SetTimestamp(timestamp)
		e.SetVersion(version)
		e.ClientID = "CLIENT_ID"

		assert.Equal("STREAM_ID", e.GetID())
		assert.Equal(timestamp, e.GetTimestamp())
		assert.Equal(version, e.GetVersion())
	})

	t.Run("account created is invalid if account id is empty", func(t *testing.T) {
		assert := require.New(t)

		var e model.AccountCreated
		err := e.Validate()

		assert.Equal(model.ErrAccountIDEmpty, err)
	})

	t.Run("account created is invalid if client id is empty", func(t *testing.T) {
		assert := require.New(t)

		var e model.AccountCreated
		e.ID = "ID"
		err := e.Validate()

		assert.Equal(model.ErrClientIDEmpty, err)
	})
}

func Test_build_account(t *testing.T) {
	t.Run("create account", func(t *testing.T) {
		assert := require.New(t)

		id := "ACCOUNT_ID"
		clientID := "CLIENT_ID"

		account := model.NewAccount(id)

		err := account.CreateAccount(clientID)

		assert.NoError(err)
		assert.Equal(id, account.GetID())
		assert.Equal(clientID, account.GetClientID())
		assert.Equal(-1, account.GetVersion())
	})

	t.Run("create account requires account id", func(t *testing.T) {
		assert := require.New(t)

		account := model.NewAccount("")

		err := account.CreateAccount("")

		assert.Equal(model.ErrAccountIDEmpty, err)
	})

	t.Run("create account requires client id with value", func(t *testing.T) {
		assert := require.New(t)

		id := "ACCOUNT_ID"

		account := model.NewAccount(id)

		err := account.CreateAccount("")

		assert.Equal(model.ErrClientIDEmpty, err)
	})

	t.Run("rebuild account", func(t *testing.T) {
		assert := require.New(t)

		id := "client_id"
		clientID := id

		account := model.NewAccount("")
		account.RebuildFrom(sampleEventsForBuild(id)...)

		assert.Equal(id, account.GetID())
		assert.Equal(clientID, account.GetClientID())
		transactions := account.GetTransactions()
		assert.Len(transactions, 3)
		assert.Equal(
			model.Transaction{
				Type:   model.DepositTransaction,
				Amount: 1000,
				Time:   parseDate("10-01-2012"),
			},
			transactions[0])
		assert.Equal(
			model.Transaction{
				Type:   model.DepositTransaction,
				Amount: 2000,
				Time:   parseDate("13-01-2012"),
			},
			transactions[1])
		assert.Equal(
			model.Transaction{
				Type:   model.WithdrawalTransaction,
				Amount: 500,
				Time:   parseDate("14-01-2012"),
			},
			transactions[2])
		assert.Equal(3, account.GetVersion())
	})

	t.Run("deposit amount", func(t *testing.T) {
		assert := require.New(t)

		var (
			id                           = "id"
			amount          model.Amount = 1000
			transactionTime              = parseDate("10-01-2012")
		)

		account := model.NewAccount(id)
		account.Deposit(amount, transactionTime)

		transactions := account.GetTransactions()
		changes := account.Changes()

		assert.Len(transactions, 1)
		assert.Len(changes, 1)
		assert.Equal(
			model.Transaction{
				Type:   model.DepositTransaction,
				Amount: 1000,
				Time:   parseDate("10-01-2012"),
			},
			transactions[0])

		var event model.AmountDeposited
		event.ID = account.GetID()
		event.Amount = amount
		event.TransactionTime = transactionTime
		assert.EqualValues(&event, changes[0])
		assert.Equal(-1, account.GetVersion())
	})

	t.Run("withdraw amount", func(t *testing.T) {
		assert := require.New(t)

		var (
			id                           = "id"
			amount          model.Amount = 1000
			transactionTime              = parseDate("10-01-2012")
		)

		account := model.NewAccount(id)
		account.Withdraw(amount, transactionTime)

		transactions := account.GetTransactions()
		changes := account.Changes()

		assert.Len(transactions, 1)
		assert.Len(changes, 1)
		assert.Equal(
			model.Transaction{
				Type:   model.WithdrawalTransaction,
				Amount: 1000,
				Time:   parseDate("10-01-2012"),
			},
			transactions[0])

		var event model.AmountWithdrawn
		event.ID = account.GetID()
		event.Amount = amount
		event.TransactionTime = transactionTime
		assert.EqualValues(&event, changes[0])
		assert.Equal(-1, account.GetVersion())
	})

	t.Run("apply events to account and get changes", func(t *testing.T) {
		assert := require.New(t)

		id := "ACCOUNT_ID"
		events := sampleEventsForBuild(id)

		account := model.NewAccount(id)
		account.RebuildFrom(events...)

		assert.Equal(id, account.GetClientID())
		transactions := account.GetTransactions()
		assert.Len(transactions, 3)
		assert.Equal(
			model.Transaction{
				Type:   model.DepositTransaction,
				Amount: 1000,
				Time:   parseDate("10-01-2012"),
			},
			transactions[0])
		assert.Equal(
			model.Transaction{
				Type:   model.DepositTransaction,
				Amount: 2000,
				Time:   parseDate("13-01-2012"),
			},
			transactions[1])
		assert.Equal(
			model.Transaction{
				Type:   model.WithdrawalTransaction,
				Amount: 500,
				Time:   parseDate("14-01-2012"),
			},
			transactions[2])

		var changes []model.StreamEvent = account.Changes()
		assert.Len(changes, 0)

		assert.Equal(3, account.GetVersion())
	})

	t.Run("call commands and get changes", func(t *testing.T) {
		assert := require.New(t)

		id := "ACCOUNT_ID"

		account := model.NewAccount(id)

		err := account.CreateAccount(id)
		assert.NoError(err)

		account.Deposit(1000, parseDate("10-01-2012"))
		account.Deposit(2000, parseDate("13-01-2012"))
		account.Withdraw(500, parseDate("14-01-2012"))

		assert.Equal(id, account.GetClientID())

		transactions := account.GetTransactions()
		assert.Len(transactions, 3)
		assert.Equal(
			model.Transaction{
				Type:   model.DepositTransaction,
				Amount: 1000,
				Time:   parseDate("10-01-2012"),
			},
			transactions[0])
		assert.Equal(
			model.Transaction{
				Type:   model.DepositTransaction,
				Amount: 2000,
				Time:   parseDate("13-01-2012"),
			},
			transactions[1])
		assert.Equal(
			model.Transaction{
				Type:   model.WithdrawalTransaction,
				Amount: 500,
				Time:   parseDate("14-01-2012"),
			},
			transactions[2])

		var changes []model.StreamEvent = account.Changes()
		assert.Len(changes, 4)
		events := sampleEventsForChanges(id)
		assert.EqualValues(events, changes)
		assert.Equal(-1, account.GetVersion())
	})

	t.Run("reload from snapshot & get snapshot", func(t *testing.T) {
		var (
			assert = require.New(t)
			id     = "ACCOUNT_ID"
		)

		var tx = model.Transaction{
			Type:   model.DepositTransaction,
			Amount: 1000,
			Time:   parseDate("10-01-2012"),
		}

		var snapshot model.AccountSnapshot
		snapshot.ID = id
		snapshot.ClientID = id
		snapshot.Version = 3
		snapshot.Transactions = []model.Transaction{tx}

		account := model.NewAccount("")
		account.ReloadFromSnapshot(&snapshot)

		assert.Equal(id, account.GetID())
		assert.Equal(id, account.GetClientID())
		assert.Equal(3, account.GetVersion())
		assert.EqualValues([]model.Transaction{tx}, account.GetTransactions())

		newSnapshot := account.GetSnapshot()
		assert.EqualValues(&snapshot, newSnapshot)
	})
}

func Test_generate_statement(t *testing.T) {
	assert := require.New(t)

	id := "ID"

	account := model.NewAccount(id)
	account.RebuildFrom(sampleEventsForBuild(id)...)

	statement := account.Statement()

	assert.Equal(expectedBankStatement, statement.String())
}

func Test_add_statement_line(t *testing.T) {
	assert := require.New(t)

	statement := model.NewStatement()

	sampleLines := sampleStatementLines()
	for _, line := range sampleLines {
		statement.AddStatementLine(line)
	}

	assert.Equal(len(sampleLines), len(statement.Lines))
}

func Test_get_bank_statement(t *testing.T) {
	assert := require.New(t)

	statement := model.NewStatement()

	sampleLines := sampleStatementLines()
	for _, line := range sampleLines {
		statement.AddStatementLine(line)
	}

	assert.Equal(expectedBankStatement, statement.String())
}

func Test_statement_line(t *testing.T) {
	assert := require.New(t)

	sampleLines := sampleStatementLines()
	for i, line := range sampleLines {
		assert.Equal(expectedStatementLineStrings[i], line.String())
	}
}

func sampleStatementLines() []model.StatementLine {
	return []model.StatementLine{
		{Date: parseDate("10-01-2012"), Credit: 1000, Debit: 0, Balance: 1000},
		{Date: parseDate("13-01-2012"), Credit: 2000, Debit: 0, Balance: 3000},
		{Date: parseDate("14-01-2012"), Credit: 0, Debit: 500, Balance: 2500},
	}
}

func sampleEventsForBuild(id string) (res []model.StreamEvent) {
	{
		e := model.AccountCreated{
			ClientID: id,
		}
		e.ID = id
		e.Version = 0
		res = append(res, &e)
	}
	{
		e := model.AmountDeposited{
			Amount: 1000,
		}
		e.ID = id
		e.TransactionTime = parseDate("10-01-2012")
		e.Version = 1
		res = append(res, &e)
	}
	{
		e := model.AmountDeposited{
			Amount: 2000,
		}
		e.ID = id
		e.TransactionTime = parseDate("13-01-2012")
		e.Version = 2
		res = append(res, &e)
	}
	{
		e := model.AmountWithdrawn{
			Amount: 500,
		}
		e.ID = id
		e.TransactionTime = parseDate("14-01-2012")
		e.Version = 3
		res = append(res, &e)
	}
	return
}

func sampleEventsForChanges(id string) (res []model.StreamEvent) {
	{
		e := model.AccountCreated{
			ClientID: id,
		}
		e.ID = id
		res = append(res, &e)
	}
	{
		e := model.AmountDeposited{
			Amount: 1000,
		}
		e.ID = id
		e.TransactionTime = parseDate("10-01-2012")
		res = append(res, &e)
	}
	{
		e := model.AmountDeposited{
			Amount: 2000,
		}
		e.ID = id
		e.TransactionTime = parseDate("13-01-2012")
		res = append(res, &e)
	}
	{
		e := model.AmountWithdrawn{
			Amount: 500,
		}
		e.ID = id
		e.TransactionTime = parseDate("14-01-2012")
		res = append(res, &e)
	}
	return
}

func parseDate(d string) time.Time {
	t, err := time.ParseInLocation("02-01-2006", d, time.UTC)
	if err != nil {
		panic(err)
	}
	return t
}

var (
	expectedBankStatement = `date || credit || debit || balance
14/01/2012 || || 500.00 || 2500.00
13/01/2012 || 2000.00 || || 3000.00
10/01/2012 || 1000.00 || || 1000.00`

	expectedStatementLineStrings = []string{
		"10/01/2012 || 1000.00 || || 1000.00",
		"13/01/2012 || 2000.00 || || 3000.00",
		"14/01/2012 || || 500.00 || 2500.00",
	}
)

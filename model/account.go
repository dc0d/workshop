package model

import (
	"errors"
	"time"
)

// Account represents a client account.
// It is assumed that each client has only one account - to keep the kata simple.
type Account struct {
	id           string
	clientID     string
	transactions []Transaction
	changes      []StreamEvent
	version      int
}

func NewAccount(id string) *Account {
	account := &Account{
		id:      id,
		version: -1,
	}

	return account
}

func (account *Account) GetID() string { return account.id }

func (account *Account) GetClientID() string { return account.clientID }

func (account *Account) GetTransactions() []Transaction {
	return append([]Transaction(nil), account.transactions...)
}

func (account *Account) GetVersion() int { return account.version }

func (account *Account) RebuildFrom(events ...StreamEvent) {
	account.applyEvents(false, events...)
}

func (account *Account) ReloadFromSnapshot(snapshot *AccountSnapshot) {
	account.id = snapshot.ID
	account.clientID = snapshot.ClientID
	account.version = snapshot.Version
	account.transactions = append([]Transaction{}, snapshot.Transactions...)
}

func (account *Account) GetSnapshot() *AccountSnapshot {
	var snapshot AccountSnapshot
	snapshot.ID = account.id
	snapshot.ClientID = account.clientID
	snapshot.Version = account.version
	snapshot.Transactions = append([]Transaction{}, account.transactions...)
	return &snapshot
}

func (account *Account) CreateAccount(clientID string) error {
	var event AccountCreated
	event.ID = account.GetID()
	event.ClientID = clientID

	return account.applyEvents(true, &event)
}

func (account *Account) Deposit(amount Amount, transactionTime time.Time) {
	var event AmountDeposited
	event.ID = account.GetID()
	event.Amount = amount
	event.TransactionTime = transactionTime

	account.applyEvents(true, &event)
}

func (account *Account) Withdraw(amount Amount, transactionTime time.Time) {
	var event AmountWithdrawn
	event.ID = account.GetID()
	event.Amount = amount
	event.TransactionTime = transactionTime

	account.applyEvents(true, &event)
}

func (account *Account) Changes() []StreamEvent {
	return account.changes
}

func (account *Account) Statement() *Statement {
	st := NewStatement()

	var balance Amount
	for _, tx := range account.transactions {
		var line StatementLine
		line.Date = tx.Time

		switch tx.Type {
		case DepositTransaction:
			line.Credit = tx.Amount
			balance += tx.Amount
		case WithdrawalTransaction:
			line.Debit = tx.Amount
			balance -= tx.Amount
		}
		line.Balance = balance

		st.AddStatementLine(line)
	}
	return st
}

func (account *Account) applyEvents(isNew bool, events ...StreamEvent) error {
	for _, e := range events {
		var tx *Transaction
		switch event := e.(type) {
		case *AccountCreated:
			if err := event.Validate(); err != nil {
				return err
			}

			account.id = event.GetID()
			account.clientID = event.ClientID
		case *AmountDeposited:
			tx = &Transaction{
				Type:   DepositTransaction,
				Amount: event.Amount,
				Time:   event.TransactionTime,
			}
		case *AmountWithdrawn:
			tx = &Transaction{
				Type:   WithdrawalTransaction,
				Amount: event.Amount,
				Time:   event.TransactionTime,
			}
		}
		if tx != nil {
			account.transactions = append(account.transactions, *tx)
		}
		if isNew {
			account.changes = append(account.changes, e)
		} else {
			account.version = e.GetVersion()
		}
	}

	return nil
}

type Transaction struct {
	Type   TransactionType
	Amount Amount
	Time   time.Time
}

type TransactionType string

type AccountSnapshot struct {
	ID           string        `json:"id,omitempty"`
	ClientID     string        `json:"client_id,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
	Version      int           `json:"version,omitempty"`
}

const (
	DepositTransaction    TransactionType = "DEPOSIT"
	WithdrawalTransaction TransactionType = "WITHDRAWAL"
)

var (
	ErrAccountIDEmpty = errors.New("account id must have value")
	ErrClientIDEmpty  = errors.New("client id must have value")
)

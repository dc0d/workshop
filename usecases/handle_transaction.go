package usecases

import (
	model "github.com/dc0d/workshop/domain_model"
)

type handleTransaction struct {
	repo model.AccountRepository
}

func NewHandleTransaction(repo model.AccountRepository) model.HandleTransaction {
	return &handleTransaction{repo: repo}
}

func (usecase *handleTransaction) Run(option model.HandleTransactionOption) error {
	var (
		clientID string
		options  model.HandleTransactionOptions
	)

	options.Apply(option)
	clientID = getClientID(options)

	account, err := usecase.repo.Find(clientID)
	if err == model.ErrAccountNotFound {
		account = model.NewAccount(clientID)
		account.CreateAccount(clientID)
	}

	sendCommand(account, options)

	return usecase.repo.Save(account)
}

func sendCommand(
	account *model.Account,
	options model.HandleTransactionOptions) {
	if options.DepositCommand != nil {
		account.Deposit(
			model.Amount(options.DepositCommand.Amount),
			options.DepositCommand.Time)
		return
	}
	account.Withdraw(
		model.Amount(options.WithdrawCommand.Amount),
		options.WithdrawCommand.Time)
}

func getClientID(options model.HandleTransactionOptions) string {
	if options.DepositCommand != nil {
		return options.DepositCommand.ClientID
	}
	return options.WithdrawCommand.ClientID
}

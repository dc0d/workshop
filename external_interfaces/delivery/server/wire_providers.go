package server

import (
	model "github.com/dc0d/workshop/domain_model"
	"github.com/dc0d/workshop/external_interfaces/infrastructure"
	"github.com/dc0d/workshop/interface_adapters/repositories"
	"github.com/dc0d/workshop/interface_adapters/web/api"
	"github.com/dc0d/workshop/usecases"

	"github.com/google/wire"
)

var (
	transactionCommandHandlerSet = wire.NewSet(
		provideAccountRepository,
		wire.Bind(new(model.AccountRepository), new(*repositories.AccountRepository)),
		provideHandleTransaction,
		wire.Bind(new(model.HandleTransaction), new(*usecases.HandleTransaction)),
		api.NewTransactionCommandHandler)

	statementHandlerSet = wire.NewSet(
		provideStatementViewRepository,
		wire.Bind(new(model.StatementViewRepository), new(*repositories.StatementViewRepository)),
		provideBankStatement,
		wire.Bind(new(model.BankStatement), new(*usecases.BankStatement)),
		api.NewStatementHandler)
)

func provideHandleTransaction(repo model.AccountRepository) *usecases.HandleTransaction {
	return usecases.NewHandleTransaction(repo)
}

func provideBankStatement(repo model.StatementViewRepository) *usecases.BankStatement {
	return usecases.NewBankStatement(repo)
}

func provideAccountRepository() *repositories.AccountRepository {
	return _accountRepo
}

func provideStatementViewRepository() *repositories.StatementViewRepository {
	return _statementRepo
}

var (
	_statementRepo = repositories.NewStatementViewRepository(_statementViewStorage)
	_accountRepo   = repositories.NewAccountRepository(_eventStore, _timeSource)
	_timeSource    = repositories.NewTimeSource()
	_eventStore    = repositories.NewEventStore(_storage, _publisher)
	_publisher     = infrastructure.NewQueueEventPublisher(_queue)

	_queue   = infrastructure.NewFakeEventQueue()
	_storage = infrastructure.NewEventStorage()

	_statementViewBuilder = infrastructure.NewStatementViewBuilder(_queue, _statementViewStorage)
	_statementViewStorage = infrastructure.NewStatementViewStorage()
)

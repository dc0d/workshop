package server

import (
	"net/http"
	"time"

	"github.com/dc0d/workshop/external_interfaces/infrastructure"
	"github.com/dc0d/workshop/interface_adapters/repositories"
	"github.com/dc0d/workshop/interface_adapters/web/api"
	"github.com/dc0d/workshop/model"
)

// Start .
func Start() {
	router := api.NewRouter(
		defaultAccountRepositoryFactory{},
		defaultStatementViewRepositoryFactory{})

	s := newServer()
	router.Logger.Fatal(router.StartServer(s))
}

func newServer() *http.Server {
	return &http.Server{
		Addr:              ":8090",
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
	}
}

type defaultAccountRepositoryFactory struct{}

func (defaultAccountRepositoryFactory) CreateAccountRepository() model.AccountRepository {
	return _accountRepo
}

type defaultStatementViewRepositoryFactory struct{}

func (defaultStatementViewRepositoryFactory) CreateStatementViewRepository() model.StatementViewRepository {
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

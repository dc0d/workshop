package server

import (
	"net/http"
	"time"

	model "github.com/dc0d/workshop/domain_model"
	"github.com/dc0d/workshop/external_interfaces/infrastructure"
	"github.com/dc0d/workshop/interface_adapters/repositories"
	"github.com/dc0d/workshop/interface_adapters/web/api"
)

func Start() {
	router := api.NewRouter(
		model.AccountRepositoryFactoryFunc(createAccountRepository),
		model.StatementViewRepositoryFactoryFunc(createStatementViewRepository))

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

func createAccountRepository() model.AccountRepository {
	return _accountRepo
}

func createStatementViewRepository() model.StatementViewRepository {
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

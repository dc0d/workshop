package api

import (
	"gitlab.com/dc0d/go-workshop/external_interfaces/infrastructure"
	"gitlab.com/dc0d/go-workshop/interface_adapters/repositories"
	"gitlab.com/dc0d/go-workshop/model"
	"gitlab.com/dc0d/go-workshop/usecases"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// NewRouter .
func NewRouter() *echo.Echo {
	return newRouter(newDefaultHandlerFactory())
}

func newRouter(handlerFactory handlerFactory) *echo.Echo {
	router := echo.New()
	router.Logger.SetLevel(log.INFO)
	router.Use(middleware.Recover())

	{
		bankAPI := router.Group("/api/bank")

		bankAPI.POST("/transactions", func(c echo.Context) error {
			handler := handlerFactory.createTransactionCommandHandler()
			return handler.handleCommand(c)
		})

		bankAPI.GET("/:client_id/statement", func(c echo.Context) error {
			handler := handlerFactory.createStatementHandler()
			return handler.getStatement(c)
		})
	}

	return router
}

type handlerFactory interface {
	createTransactionCommandHandler() *transactionCommandHandler
	createStatementHandler() *statementHandler
}

type defaultHandlerFactory struct{}

func newDefaultHandlerFactory() *defaultHandlerFactory { return &defaultHandlerFactory{} }

func (*defaultHandlerFactory) createTransactionCommandHandler() *transactionCommandHandler {
	repo := createAccountRepository()
	usecase := usecases.NewHandleTransaction(repo)
	return newTransactionCommandHandler(usecase)
}

func (*defaultHandlerFactory) createStatementHandler() *statementHandler {
	repo := createStatementViewRepository()
	usecase := usecases.NewBankStatement(repo)
	return newStatementHandler(usecase)
}

func createStatementViewRepository() model.StatementViewRepository {
	return _statementRepo
}

func createAccountRepository() model.AccountRepository {
	return _accountRepo
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

package api

import (
	"github.com/dc0d/workshop/model"
	"github.com/dc0d/workshop/usecases"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

// NewRouter .
func NewRouter(
	accountRepositoryFactory model.AccountRepositoryFactory,
	statementViewRepositoryFactory model.StatementViewRepositoryFactory) *echo.Echo {
	return newRouter(newDefaultHandlerFactory(accountRepositoryFactory, statementViewRepositoryFactory))
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

type defaultHandlerFactory struct {
	accountRepositoryFactory       model.AccountRepositoryFactory
	statementViewRepositoryFactory model.StatementViewRepositoryFactory
}

func newDefaultHandlerFactory(
	accountRepositoryFactory model.AccountRepositoryFactory,
	statementViewRepositoryFactory model.StatementViewRepositoryFactory) *defaultHandlerFactory {
	return &defaultHandlerFactory{
		accountRepositoryFactory:       accountRepositoryFactory,
		statementViewRepositoryFactory: statementViewRepositoryFactory,
	}
}

func (factory *defaultHandlerFactory) createTransactionCommandHandler() *transactionCommandHandler {
	repo := factory.accountRepositoryFactory.CreateAccountRepository()
	usecase := usecases.NewHandleTransaction(repo)
	return newTransactionCommandHandler(usecase)
}

func (factory *defaultHandlerFactory) createStatementHandler() *statementHandler {
	repo := factory.statementViewRepositoryFactory.CreateStatementViewRepository()
	usecase := usecases.NewBankStatement(repo)
	return newStatementHandler(usecase)
}

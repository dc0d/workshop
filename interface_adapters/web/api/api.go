package api

//go:generate moq -pkg api -out ./moq_handle_transaction_test.go ./../../../domain_model HandleTransaction
//go:generate moq -pkg api -out ./moq_handle_transaction_factory_test.go ./../../../domain_model HandleTransactionFactory
//go:generate moq -pkg api -out ./moq_account_repository_factory_test.go ./../../../domain_model AccountRepositoryFactory

//go:generate moq -pkg api -out ./moq_bank_statement_test.go ./../../../domain_model BankStatement
//go:generate moq -pkg api -out ./moq_bank_statement_factory_test.go ./../../../domain_model BankStatementFactory
//go:generate moq -pkg api -out ./moq_statement_view_repository_factory_test.go ./../../../domain_model StatementViewRepositoryFactory

import (
	model "github.com/dc0d/workshop/domain_model"
	"github.com/dc0d/workshop/usecases"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(
	accountRepositoryFactory model.AccountRepositoryFactory,
	statementViewRepositoryFactory model.StatementViewRepositoryFactory) *echo.Echo {
	var ctxFac ProviderContextFactory = func(c echo.Context) ProviderContext {
		return newAppCtx(c, accountRepositoryFactory, statementViewRepositoryFactory)
	}
	return newRouter(ctxFac)
}

type appCtx struct {
	model.AccountRepositoryFactory
	model.StatementViewRepositoryFactory

	usecaseTx  model.HandleTransaction
	usecaseStt model.BankStatement

	echo.Context
}

func newAppCtx(
	c echo.Context,
	accountRepositoryFactory model.AccountRepositoryFactory,
	statementViewRepositoryFactory model.StatementViewRepositoryFactory) *appCtx {
	res := &appCtx{
		AccountRepositoryFactory:       accountRepositoryFactory,
		StatementViewRepositoryFactory: statementViewRepositoryFactory,
		Context:                        c,
	}

	return res
}

func (ctx *appCtx) CreateHandleTransaction(repo model.AccountRepository) model.HandleTransaction {
	if ctx.usecaseTx == nil {
		ctx.usecaseTx = usecases.NewHandleTransaction(repo)
	}

	return ctx.usecaseTx
}

func (ctx *appCtx) CreateBankStatement(repo model.StatementViewRepository) model.BankStatement {
	if ctx.usecaseStt == nil {
		ctx.usecaseStt = usecases.NewBankStatement(repo)
	}

	return ctx.usecaseStt
}

func newRouter(ctxFac ProviderContextFactory) *echo.Echo {
	router := echo.New()

	router.Use(middleware.Recover())

	{
		bankAPI := router.Group("/api/bank")
		bankAPI.Use(providers(ctxFac))

		bankAPI.POST("/transactions", func(c echo.Context) error {
			var ctx = c.(ProviderContext)

			repo := ctx.CreateAccountRepository()
			usecase := ctx.CreateHandleTransaction(repo)
			handler := newTransactionCommandHandler(usecase)

			return handler.handleCommand(ctx)
		})

		bankAPI.GET("/:client_id/statement", func(c echo.Context) error {
			var ctx = c.(ProviderContext)

			repo := ctx.CreateStatementViewRepository()
			usecase := ctx.CreateBankStatement(repo)
			handler := newStatementHandler(usecase)

			return handler.getStatement(ctx)
		})
	}

	return router
}

func providers(ctxFac ProviderContextFactory) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c = ctxFac(c)
			return next(c)
		}
	}
}

type ProviderContextFactory func(echo.Context) ProviderContext

type ProviderContext interface {
	model.AccountRepositoryFactory
	model.StatementViewRepositoryFactory
	model.HandleTransactionFactory
	model.BankStatementFactory

	echo.Context
}

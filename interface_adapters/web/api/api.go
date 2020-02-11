package api

//go:generate moq -pkg api -out ./moq_handle_transaction_test.go ./../../../domain_model HandleTransaction
//go:generate moq -pkg api -out ./moq_bank_statement_test.go ./../../../domain_model BankStatement

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter(injectors InjectorSet) *echo.Echo {
	router := echo.New()

	router.Use(middleware.Recover())

	{
		bankAPI := router.Group("/api/bank")

		bankAPI.POST("/transactions", func(c echo.Context) error {
			handler := injectors.InjectTransactionCommandHandler()

			return handler.handleCommand(c)
		})

		bankAPI.GET("/:client_id/statement", func(c echo.Context) error {
			handler := injectors.InjectStatementHandler()

			return handler.getStatement(c)
		})
	}

	return router
}

type InjectorSet struct {
	InjectStatementHandler          func() *StatementHandler
	InjectTransactionCommandHandler func() *TransactionCommandHandler
}

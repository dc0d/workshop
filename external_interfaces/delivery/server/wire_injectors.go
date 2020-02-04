// +build wireinject

package server

import (
	"github.com/dc0d/workshop/interface_adapters/web/api"

	"github.com/google/wire"
)

func injectStatementHandler() (res *api.StatementHandler) {
	wire.Build(statementHandlerSet)
	return
}

func injectTransactionCommandHandler() (res *api.TransactionCommandHandler) {
	wire.Build(transactionCommandHandlerSet)
	return
}

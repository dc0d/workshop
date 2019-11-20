package api

import (
	"net/http"

	"github.com/dc0d/workshop/model"

	"github.com/labstack/echo"
)

type statementHandler struct {
	usecase model.BankStatement
}

func newStatementHandler(usecase model.BankStatement) *statementHandler {
	return &statementHandler{usecase: usecase}
}

func (h *statementHandler) getStatement(c echo.Context) error {
	statement, err := h.usecase.Run(c.Param("client_id"))
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, statement.String())
}

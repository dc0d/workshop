package api

import (
	"net/http"

	model "github.com/dc0d/workshop/domain_model"

	"github.com/labstack/echo/v4"
)

type StatementHandler struct {
	usecase model.BankStatement
}

func NewStatementHandler(usecase model.BankStatement) *StatementHandler {
	return &StatementHandler{usecase: usecase}
}

func (h *StatementHandler) getStatement(c echo.Context) error {
	statement, err := h.usecase.Run(c.Param("client_id"))
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, statement.String())
}

package api

import (
	"errors"
	"net/http"
	"time"

	model "github.com/dc0d/workshop/domain_model"
	"github.com/labstack/echo/v4"
)

type TransactionCommandHandler struct {
	usecase model.HandleTransaction
}

func NewTransactionCommandHandler(usecase model.HandleTransaction) *TransactionCommandHandler {
	return &TransactionCommandHandler{usecase: usecase}
}

func (h *TransactionCommandHandler) handleCommand(c echo.Context) error {
	var command transactionCommand
	if err := c.Bind(&command); err != nil {
		return err
	}

	usecaseOption, err := transactionCommandToUsecaseOption(command)
	if err != nil {
		return err
	}

	err = h.usecase.Run(usecaseOption)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func transactionCommandToUsecaseOption(command transactionCommand) (model.HandleTransactionOption, error) {
	var (
		clientID        = command.Data.ClientID
		commandType     = command.Command
		amount          = command.Data.Amount
		transactionTime = command.Data.Time
	)

	switch commandType {
	case depositCommandName:
		return model.DepositWith(
			model.DepositCommand{
				ClientID: clientID,
				Amount:   amount,
				Time:     transactionTime,
			}), nil
	case withdrawCommandName:
		return model.WithdrawWith(
			model.WithdrawCommand{
				ClientID: clientID,
				Amount:   amount,
				Time:     transactionTime,
			}), nil
	}

	return nil, errUnknownCommandType
}

type transactionCommand struct {
	Command string `json:"command"`
	Data    struct {
		ClientID string    `json:"client_id"`
		Amount   int       `json:"amount"`
		Time     time.Time `json:"time"`
	} `json:"data"`
}

var (
	errUnknownCommandType = errors.New("unknown command type")
)

const (
	depositCommandName  = "DEPOSIT"
	withdrawCommandName = "WITHDRAW"
)

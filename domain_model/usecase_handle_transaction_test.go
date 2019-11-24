package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_handle_transaction_options(t *testing.T) {
	t.Run("deposit with", func(t *testing.T) {
		assert := require.New(t)

		depositCommand := DepositCommand{
			ClientID: "ID",
			Amount:   10000000000,
			Time:     time.Now(),
		}

		var options HandleTransactionOptions

		options.Apply(DepositWith(depositCommand))

		assert.Equal(depositCommand, *options.DepositCommand)
	})

	t.Run("deposit with sets withdraw command to nil", func(t *testing.T) {
		assert := require.New(t)

		depositCommand := DepositCommand{
			ClientID: "ID",
			Amount:   10000000000,
			Time:     time.Now(),
		}

		var options HandleTransactionOptions
		options.WithdrawCommand = &WithdrawCommand{}

		options.Apply(DepositWith(depositCommand))

		assert.Equal(depositCommand, *options.DepositCommand)
		assert.Nil(options.WithdrawCommand)
	})

	t.Run("withdraw with", func(t *testing.T) {
		assert := require.New(t)

		withdrawCommand := WithdrawCommand{
			ClientID: "ID",
			Amount:   10000000000,
			Time:     time.Now(),
		}

		var options HandleTransactionOptions
		options.DepositCommand = &DepositCommand{}

		options.Apply(WithdrawWith(withdrawCommand))

		assert.Equal(withdrawCommand, *options.WithdrawCommand)
		assert.Nil(options.DepositCommand)
	})
}

package model

import (
	"time"
)

type HandleTransaction interface {
	Run(HandleTransactionOption) error
}

type DepositCommand struct {
	ClientID string
	Amount   int
	Time     time.Time
}

type WithdrawCommand struct {
	ClientID string
	Amount   int
	Time     time.Time
}

type HandleTransactionOptions struct {
	DepositCommand  *DepositCommand
	WithdrawCommand *WithdrawCommand
}

func (options *HandleTransactionOptions) Apply(list ...HandleTransactionOption) {
	for _, f := range list {
		f(options)
	}
}

type HandleTransactionOption func(*HandleTransactionOptions)

func DepositWith(command DepositCommand) HandleTransactionOption {
	return func(opt *HandleTransactionOptions) {
		opt.DepositCommand = &command
		opt.WithdrawCommand = nil
	}
}

func WithdrawWith(command WithdrawCommand) HandleTransactionOption {
	return func(opt *HandleTransactionOptions) {
		opt.DepositCommand = nil
		opt.WithdrawCommand = &command
	}
}

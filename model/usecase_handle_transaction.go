package model

import (
	"time"
)

// HandleTransaction .
type HandleTransaction interface {
	// Run .
	Run(HandleTransactionOption) error
}

// DepositCommand .
type DepositCommand struct {
	ClientID string
	Amount   int
	Time     time.Time
}

// WithdrawCommand .
type WithdrawCommand struct {
	ClientID string
	Amount   int
	Time     time.Time
}

// HandleTransactionOptions .
type HandleTransactionOptions struct {
	DepositCommand  *DepositCommand
	WithdrawCommand *WithdrawCommand
}

// Apply .
func (options *HandleTransactionOptions) Apply(list ...HandleTransactionOption) {
	for _, f := range list {
		f(options)
	}
}

// HandleTransactionOption .
type HandleTransactionOption func(*HandleTransactionOptions)

// DepositWith .
func DepositWith(command DepositCommand) HandleTransactionOption {
	return func(opt *HandleTransactionOptions) {
		opt.DepositCommand = &command
		opt.WithdrawCommand = nil
	}
}

// WithdrawWith .
func WithdrawWith(command WithdrawCommand) HandleTransactionOption {
	return func(opt *HandleTransactionOptions) {
		opt.DepositCommand = nil
		opt.WithdrawCommand = &command
	}
}

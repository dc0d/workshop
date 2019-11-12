package model

// BankStatement .
type BankStatement interface {
	Run(id string) (*Statement, error)
}

package model

type BankStatement interface {
	Run(id string) (*Statement, error)
}

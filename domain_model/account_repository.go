package model

type AccountRepository interface {
	Find(string) (*Account, error)
	Save(*Account) error
}

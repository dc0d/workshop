package model

// StatementViewRepositoryFactory .
type StatementViewRepositoryFactory interface {
	CreateStatementViewRepository() StatementViewRepository
}

// AccountRepositoryFactory .
type AccountRepositoryFactory interface {
	CreateAccountRepository() AccountRepository
}

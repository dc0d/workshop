package model

type StatementViewRepositoryFactory interface {
	CreateStatementViewRepository() StatementViewRepository
}

type AccountRepositoryFactory interface {
	CreateAccountRepository() AccountRepository
}

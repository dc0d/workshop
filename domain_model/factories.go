package model

type StatementViewRepositoryFactory interface {
	CreateStatementViewRepository() StatementViewRepository
}

type AccountRepositoryFactory interface {
	CreateAccountRepository() AccountRepository
}

type StatementViewRepositoryFactoryFunc func() StatementViewRepository

func (f StatementViewRepositoryFactoryFunc) CreateStatementViewRepository() StatementViewRepository {
	return f()
}

type AccountRepositoryFactoryFunc func() AccountRepository

func (f AccountRepositoryFactoryFunc) CreateAccountRepository() AccountRepository {
	return f()
}

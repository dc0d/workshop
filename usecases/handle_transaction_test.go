package usecases_test

import (
	"testing"

	model "github.com/dc0d/workshop/domain_model"
	"github.com/dc0d/workshop/usecases"

	"github.com/stretchr/testify/require"
)

func Test_create_usecase(t *testing.T) {
	var _ model.HandleTransaction = usecases.NewHandleTransaction(nil)
}

func Test_deposit_command(t *testing.T) {
	assert := require.New(t)

	var (
		id         = "A_CLIENT_ID"
		findCalled int
		saveCalled int
	)

	repo := newRepoSpy(
		func(_id string) (*model.Account, error) {
			findCalled++
			assert.Equal(id, _id)
			return nil, model.ErrAccountNotFound
		},
		func(account *model.Account) error {
			saveCalled++
			assert.Equal(2, len(account.Changes()))
			return nil
		})

	command := model.DepositCommand{
		ClientID: id,
		Amount:   1000,
		Time:     parseDate("02-11-2019"),
	}

	usecase := usecases.NewHandleTransaction(repo)
	err := usecase.Run(model.DepositWith(command))

	assert.NoError(err)
	assert.Equal(1, findCalled)
	assert.Equal(1, saveCalled)
}

func Test_withdraw_command(t *testing.T) {
	assert := require.New(t)

	var (
		id         = "A_CLIENT_ID"
		findCalled int
		saveCalled int
	)

	repo := newRepoSpy(
		func(_id string) (*model.Account, error) {
			findCalled++
			assert.Equal(id, _id)
			return nil, model.ErrAccountNotFound
		},
		func(account *model.Account) error {
			saveCalled++
			assert.Equal(2, len(account.Changes()))
			return nil
		})

	command := model.WithdrawCommand{
		ClientID: id,
		Amount:   1000,
		Time:     parseDate("02-11-2019"),
	}

	usecase := usecases.NewHandleTransaction(repo)
	err := usecase.Run(model.WithdrawWith(command))

	assert.NoError(err)
	assert.Equal(1, findCalled)
	assert.Equal(1, saveCalled)
}

func Test_deposit_command_existing_account(t *testing.T) {
	assert := require.New(t)

	var (
		id         = "A_CLIENT_ID"
		findCalled int
		saveCalled int
	)

	repo := newRepoSpy(
		func(_id string) (*model.Account, error) {
			findCalled++
			assert.Equal(id, _id)
			account := model.NewAccount("")
			account.RebuildFrom(sampleEventsForBuild(id)...)
			return account, nil
		},
		func(account *model.Account) error {
			saveCalled++
			changes := account.Changes()
			assert.Equal(1, len(changes))

			event := changes[0].(*model.AmountDeposited)
			assert.Equal("A_CLIENT_ID", event.GetID())
			assert.Equal(parseDate("01-12-2019"), event.TransactionTime)
			assert.Equal(model.Amount(10000000000), event.Amount)

			return nil
		})

	command := model.DepositCommand{
		ClientID: id,
		Amount:   10000000000,
		Time:     parseDate("01-12-2019"),
	}

	usecase := usecases.NewHandleTransaction(repo)
	err := usecase.Run(model.DepositWith(command))

	assert.NoError(err)
	assert.Equal(1, findCalled)
	assert.Equal(1, saveCalled)
}

var _ model.AccountRepository = newRepoSpy(nil, nil)

type repoSpy struct {
	onFind func(string) (*model.Account, error)
	onSave func(*model.Account) error
}

func newRepoSpy(
	onFind func(string) (*model.Account, error),
	onSave func(*model.Account) error) *repoSpy {
	return &repoSpy{
		onFind: onFind,
		onSave: onSave,
	}
}

func (repo *repoSpy) Find(id string) (*model.Account, error) { return repo.onFind(id) }
func (repo *repoSpy) Save(account *model.Account) error      { return repo.onSave(account) }

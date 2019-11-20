package repositories_test

//go:generate moq -pkg repositories_test -out ./mock_statement_view_storage_read_test.go . StatementViewStorageRead

import (
	"testing"
	"time"

	"github.com/dc0d/workshop/interface_adapters/repositories"
	"github.com/dc0d/workshop/model"

	"github.com/stretchr/testify/require"
)

func Test_create_statement_view_repository(t *testing.T) {
	var view repositories.StatementViewStorageRead
	var _ model.StatementViewRepository = repositories.NewStatementViewRepository(view)
}

func Test_statement_view_repository_find(t *testing.T) {
	var (
		assert            = require.New(t)
		id                = "ID"
		now               = time.Now().UTC()
		expectedStatement = model.Statement{
			Lines: []model.StatementLine{
				{
					Date:    now,
					Credit:  1000,
					Debit:   1000,
					Balance: 1000,
				},
			},
		}
	)

	view := &StatementViewStorageReadMock{
		FindFunc: func(id string) (*model.Statement, error) {
			return &model.Statement{
				Lines: []model.StatementLine{
					{
						Date:    now,
						Credit:  1000,
						Debit:   1000,
						Balance: 1000,
					},
				},
			}, nil
		},
	}
	repo := repositories.NewStatementViewRepository(view)

	found, _ := repo.Find(id)

	assert.Len(view.FindCalls(), 1)
	assert.EqualValues(&expectedStatement, found)
}

package infrastructure_test

import (
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/dc0d/workshop/external_interfaces/infrastructure"
	"github.com/dc0d/workshop/model"

	"github.com/stretchr/testify/require"
)

func Test_creat_statement_view_storage(t *testing.T) {
	var _ model.StatementViewStorage = infrastructure.NewStatementViewStorage()
}

func Test_statement_view_storage_save_find(t *testing.T) {
	var (
		assert = require.New(t)
		id     = "ID"
	)

	storage := infrastructure.NewStatementViewStorage()

	{
		accountEvents := sampleEventsForBuild(id)
		var eventRecords []model.EventRecord
		for _, e := range accountEvents {
			var rec model.EventRecord
			rec.StreamID = e.GetID()
			rec.Version = e.GetVersion()
			rec.Data = toEventRecordData(e)

			eventRecords = append(eventRecords, rec)
		}

		err := storage.Save(eventRecords...)
		assert.NoError(err)
	}

	{
		statement, err := storage.Find(id)
		assert.NoError(err)
		assert.Equal(expectedBankStatement, statement.String())
	}
}

func Test_account_not_found(t *testing.T) {
	var (
		assert = require.New(t)
		id     = "ID"
	)

	storage := infrastructure.NewStatementViewStorage()

	_, err := storage.Find(id)
	assert.EqualValues(model.ErrAccountNotFound, err)
}

func toEventRecordData(event interface{}) []byte {
	data := toJSON(event)

	var rec model.EventRecordData
	rec.Type = typeOf(event)
	rec.EventData = data

	data = toJSON(rec)

	return data
}

func typeOf(v interface{}) string {
	parts := strings.Split(reflect.TypeOf(v).String(), ".")
	return parts[len(parts)-1]
}

func sampleEventsForBuild(id string) (res []model.StreamEvent) {
	{
		e := model.AccountCreated{
			ClientID: id,
		}
		e.ID = id
		e.Timestamp = parseDate("01-01-2019")
		e.Version = 0
		res = append(res, &e)
	}
	{
		e := model.AmountDeposited{
			Amount: 1000,
		}
		e.ID = id
		e.Timestamp = parseDate("10-01-2019")
		e.TransactionTime = parseDate("10-01-2012")
		e.Version = 1
		res = append(res, &e)
	}
	{
		e := model.AmountDeposited{
			Amount: 2000,
		}
		e.ID = id
		e.Timestamp = parseDate("13-01-2019")
		e.TransactionTime = parseDate("13-01-2012")
		e.Version = 2
		res = append(res, &e)
	}
	{
		e := model.AmountWithdrawn{
			Amount: 500,
		}
		e.ID = id
		e.Timestamp = parseDate("14-01-2019")
		e.TransactionTime = parseDate("14-01-2012")
		e.Version = 3
		res = append(res, &e)
	}
	return
}

func parseDate(d string) time.Time {
	t, err := time.ParseInLocation("02-01-2006", d, time.UTC)
	if err != nil {
		panic(err)
	}
	return t
}

var (
	expectedBankStatement = `date || credit || debit || balance
14/01/2012 || || 500.00 || 2500.00
13/01/2012 || 2000.00 || || 3000.00
10/01/2012 || 1000.00 || || 1000.00`
)

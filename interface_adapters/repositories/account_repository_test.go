package repositories_test

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/dc0d/workshop/interface_adapters/repositories"
	model "github.com/dc0d/workshop/domain_model"

	"github.com/stretchr/testify/require"
)

func Test_account_repository_interface(t *testing.T) {
	var _ model.AccountRepository = repositories.NewAccountRepository(nil, nil)
}

func Test_save_account(t *testing.T) {
	assert := require.New(t)

	id := "CLIENT_ID"
	account := model.NewAccount(id)
	account.CreateAccount(id)
	account.Deposit(1000, parseDate("10-01-2012"))
	account.Deposit(2000, parseDate("13-01-2012"))
	account.Withdraw(500, parseDate("14-01-2012"))

	var appendCalled bool
	store := newESSpy(func(events ...model.EventRecord) error {
		appendCalled = true
		assert.True(len(events) > 0)

		samples := sampleEventsForBuild("CLIENT_ID")
		for i, e := range events {
			event := convertEventRecordDataToEvent(e.Data)
			assert.EqualValues(samples[i], event)
		}

		return nil
	}, nil)

	eventTimes := []time.Time{
		parseDate("01-01-2019"),
		parseDate("10-01-2019"),
		parseDate("13-01-2019"),
		parseDate("14-01-2019"),
	}

	nowUTC := func() time.Time {
		t := eventTimes[0]
		if len(eventTimes) > 0 {
			eventTimes = eventTimes[1:]
		}
		return t
	}

	ts := newTSSpy(nowUTC)

	repo := repositories.NewAccountRepository(store, ts)
	assert.NoError(repo.Save(account))
	assert.True(appendCalled)
}

func Test_find_account(t *testing.T) {
	assert := require.New(t)

	id := "CLIENT_ID"

	var loadCalled bool
	store := newESSpy(nil, func(streamID string) ([]model.EventRecord, error) {
		loadCalled = true
		return expectedRecords(), nil
	})

	repo := repositories.NewAccountRepository(store, nil)
	account, err := repo.Find(id)
	assert.NoError(err)
	assert.True(loadCalled)

	assert.Equal(id, account.GetID())
	assert.Equal(3, account.GetVersion())
	assert.Equal(id, account.GetClientID())

	statement := account.Statement()
	assert.Equal(expectedBankStatement, statement.String())
}

func Test_account_not_found(t *testing.T) {
	assert := require.New(t)

	id := "CLIENT_ID"

	var loadCalled bool
	store := newESSpy(nil, func(streamID string) ([]model.EventRecord, error) {
		loadCalled = true
		return nil, model.ErrStreamNotFound
	})

	repo := repositories.NewAccountRepository(store, nil)
	account, err := repo.Find(id)
	assert.Nil(account)
	assert.Equal(model.ErrAccountNotFound, err)
	assert.True(loadCalled)

}

func convertEventRecordDataToEvent(data []byte) interface{} {
	var rec eventRecordData
	fromJSON(data, &rec)
	switch rec.Type {
	case "AccountCreated":
		var e model.AccountCreated
		fromJSON(rec.EventData, &e)
		return &e
	case "AmountDeposited":
		var e model.AmountDeposited
		fromJSON(rec.EventData, &e)
		return &e
	case "AmountWithdrawn":
		var e model.AmountWithdrawn
		fromJSON(rec.EventData, &e)
		return &e
	}
	return nil
}

func fromJSON(data []byte, v interface{}) {
	if err := json.Unmarshal(data, v); err != nil {
		panic(err)
	}
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

func expectedRecords() []model.EventRecord {
	data := `[
{
	"StreamID": "CLIENT_ID",
	"Version": 0,
	"Data": "eyJUeXBlIjoiQWNjb3VudENyZWF0ZWQiLCJFdmVudERhdGEiOiJleUpKUkNJNklrTk1TVVZPVkY5SlJDSXNJbFJwYldWemRHRnRjQ0k2SWpJd01Ua3RNREV0TURGVU1EQTZNREE2TURCYUlpd2lWbVZ5YzJsdmJpSTZNQ3dpUTJ4cFpXNTBTVVFpT2lKRFRFbEZUbFJmU1VRaWZRPT0ifQ=="
},
{
	"StreamID": "CLIENT_ID",
	"Version": 1,
	"Data": "eyJUeXBlIjoiQW1vdW50RGVwb3NpdGVkIiwiRXZlbnREYXRhIjoiZXlKSlJDSTZJa05NU1VWT1ZGOUpSQ0lzSWxScGJXVnpkR0Z0Y0NJNklqSXdNVGt0TURFdE1UQlVNREE2TURBNk1EQmFJaXdpVm1WeWMybHZiaUk2TVN3aVFXMXZkVzUwSWpveE1EQXdMQ0pVY21GdWMyRmpkR2x2YmxScGJXVWlPaUl5TURFeUxUQXhMVEV3VkRBd09qQXdPakF3V2lKOSJ9"
},
{
	"StreamID": "CLIENT_ID",
	"Version": 2,
	"Data": "eyJUeXBlIjoiQW1vdW50RGVwb3NpdGVkIiwiRXZlbnREYXRhIjoiZXlKSlJDSTZJa05NU1VWT1ZGOUpSQ0lzSWxScGJXVnpkR0Z0Y0NJNklqSXdNVGt0TURFdE1UTlVNREE2TURBNk1EQmFJaXdpVm1WeWMybHZiaUk2TWl3aVFXMXZkVzUwSWpveU1EQXdMQ0pVY21GdWMyRmpkR2x2YmxScGJXVWlPaUl5TURFeUxUQXhMVEV6VkRBd09qQXdPakF3V2lKOSJ9"
},
{
	"StreamID": "CLIENT_ID",
	"Version": 3,
	"Data": "eyJUeXBlIjoiQW1vdW50V2l0aGRyYXduIiwiRXZlbnREYXRhIjoiZXlKSlJDSTZJa05NU1VWT1ZGOUpSQ0lzSWxScGJXVnpkR0Z0Y0NJNklqSXdNVGt0TURFdE1UUlVNREE2TURBNk1EQmFJaXdpVm1WeWMybHZiaUk2TXl3aVFXMXZkVzUwSWpvMU1EQXNJbFJ5WVc1ellXTjBhVzl1VkdsdFpTSTZJakl3TVRJdE1ERXRNVFJVTURBNk1EQTZNREJhSW4wPSJ9"
}
]`
	var res []model.EventRecord
	if err := json.Unmarshal([]byte(data), &res); err != nil {
		panic(err)
	}
	return res
}

func parseDate(d string) time.Time {
	t, err := time.ParseInLocation("02-01-2006", d, time.UTC)
	if err != nil {
		panic(err)
	}
	return t
}

type eventRecordData struct {
	Type      string
	EventData []byte
}

var (
	expectedBankStatement = `date || credit || debit || balance
14/01/2012 || || 500.00 || 2500.00
13/01/2012 || 2000.00 || || 3000.00
10/01/2012 || 1000.00 || || 1000.00`
)

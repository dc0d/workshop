package repositories

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/dc0d/workshop/model"
)

type AccountRepository struct {
	store      model.EventStorage
	timeSource model.TimeSource
}

func NewAccountRepository(store model.EventStorage, timeSource model.TimeSource) *AccountRepository {
	return &AccountRepository{
		store:      store,
		timeSource: timeSource,
	}
}

func (repo *AccountRepository) Find(accountID string) (*model.Account, error) {
	records, err := repo.store.Load(accountID)
	if err != nil {
		if err == model.ErrStreamNotFound {
			err = model.ErrAccountNotFound
		}
		return nil, err
	}
	var events []model.StreamEvent
	for _, rec := range records {
		e, err := convertEventRecordDataToEvent(rec.Data)
		if err != nil {
			return nil, err
		}
		events = append(events, e)
	}
	account := model.NewAccount(accountID)
	account.RebuildFrom(events...)

	return account, nil
}

func (repo *AccountRepository) Save(account *model.Account) (err error) {
	var eventRecords []model.EventRecord
	version := account.GetVersion()
	for _, e := range account.Changes() {
		version++
		var record model.EventRecord
		record.StreamID = e.GetID()
		record.Version = version

		e.SetVersion(version)
		e.SetTimestamp(repo.timeSource.NowUTC())

		data, err := toEventRecordData(e)
		if err != nil {
			return err
		}

		record.Data = data

		eventRecords = append(eventRecords, record)
	}
	return repo.store.Append(eventRecords...)
}

func convertEventRecordDataToEvent(data []byte) (model.StreamEvent, error) {
	var rec eventRecordData
	if err := fromJSON(data, &rec); err != nil {
		return nil, err
	}
	switch rec.Type {
	case "AccountCreated":
		var e model.AccountCreated
		if err := fromJSON(rec.EventData, &e); err != nil {
			return nil, err
		}
		return &e, nil
	case "AmountDeposited":
		var e model.AmountDeposited
		if err := fromJSON(rec.EventData, &e); err != nil {
			return nil, err
		}
		return &e, nil
	case "AmountWithdrawn":
		var e model.AmountWithdrawn
		if err := fromJSON(rec.EventData, &e); err != nil {
			return nil, err
		}
		return &e, nil
	}
	return nil, nil
}

func fromJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func toEventRecordData(event interface{}) ([]byte, error) {
	data, err := toJSON(event)
	if err != nil {
		return nil, err
	}

	var rec eventRecordData
	rec.Type = typeOf(event)
	rec.EventData = data

	data, err = toJSON(rec)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func toJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func typeOf(v interface{}) string {
	parts := strings.Split(reflect.TypeOf(v).String(), ".")
	return parts[len(parts)-1]
}

type eventRecordData = model.EventRecordData

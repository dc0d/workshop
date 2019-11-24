package infrastructure

import (
	"encoding/json"
	"sync"

	model "github.com/dc0d/workshop/domain_model"
)

type StatementViewStorage struct {
	mx      sync.RWMutex
	storage map[string]string
}

func NewStatementViewStorage() *StatementViewStorage {
	return &StatementViewStorage{
		storage: make(map[string]string),
	}
}

func (view *StatementViewStorage) Find(id string) (*model.Statement, error) {
	view.mx.RLock()
	defer view.mx.RUnlock()

	snapshot := view.findAccount(id)
	if snapshot == nil {
		return nil, model.ErrAccountNotFound
	}
	account := model.NewAccount("")
	account.ReloadFromSnapshot(snapshot)
	return account.Statement(), nil
}

func (view *StatementViewStorage) Save(eventRecords ...model.EventRecord) error {
	view.mx.Lock()
	defer view.mx.Unlock()

	for _, rec := range eventRecords {
		e, err := convertEventRecordDataToEvent(rec.Data)
		if err != nil {
			return err
		}
		accountID := e.GetID()

		var account *model.Account

		snapshot := view.findAccount(accountID)
		if snapshot != nil {
			account = model.NewAccount("")
			account.ReloadFromSnapshot(snapshot)
		} else {
			account = model.NewAccount(accountID)
		}

		account.RebuildFrom(e)
		snapshot = account.GetSnapshot()
		js, err := toJSON(snapshot)
		if err != nil {
			return err
		}
		view.storage[accountID] = string(js)
	}

	return nil
}

func (view *StatementViewStorage) findAccount(id string) (account *model.AccountSnapshot) {
	js, ok := view.storage[id]
	if !ok {
		return nil
	}
	var data model.AccountSnapshot
	err := fromJSON([]byte(js), &data)
	if err != nil {
		delete(view.storage, id)
		return nil
	}
	return &data
}

func convertEventRecordDataToEvent(data []byte) (model.StreamEvent, error) {
	var rec model.EventRecordData
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

func toJSON(v interface{}) ([]byte, error) {
	js, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return js, nil
}

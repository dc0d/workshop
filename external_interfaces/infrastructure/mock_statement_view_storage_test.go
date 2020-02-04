// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package infrastructure_test

import (
	"github.com/dc0d/workshop/domain_model"
	"sync"
)

var (
	lockStatementViewStorageMockFind sync.RWMutex
	lockStatementViewStorageMockSave sync.RWMutex
)

// Ensure, that StatementViewStorageMock does implement model.StatementViewStorage.
// If this is not the case, regenerate this file with moq.
var _ model.StatementViewStorage = &StatementViewStorageMock{}

// StatementViewStorageMock is a mock implementation of model.StatementViewStorage.
//
//     func TestSomethingThatUsesStatementViewStorage(t *testing.T) {
//
//         // make and configure a mocked model.StatementViewStorage
//         mockedStatementViewStorage := &StatementViewStorageMock{
//             FindFunc: func(id string) (*model.Statement, error) {
// 	               panic("mock out the Find method")
//             },
//             SaveFunc: func(in1 ...model.EventRecord) error {
// 	               panic("mock out the Save method")
//             },
//         }
//
//         // use mockedStatementViewStorage in code that requires model.StatementViewStorage
//         // and then make assertions.
//
//     }
type StatementViewStorageMock struct {
	// FindFunc mocks the Find method.
	FindFunc func(id string) (*model.Statement, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(in1 ...model.EventRecord) error

	// calls tracks calls to the methods.
	calls struct {
		// Find holds details about calls to the Find method.
		Find []struct {
			// ID is the id argument value.
			ID string
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// In1 is the in1 argument value.
			In1 []model.EventRecord
		}
	}
}

// Find calls FindFunc.
func (mock *StatementViewStorageMock) Find(id string) (*model.Statement, error) {
	if mock.FindFunc == nil {
		panic("StatementViewStorageMock.FindFunc: method is nil but StatementViewStorage.Find was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	lockStatementViewStorageMockFind.Lock()
	mock.calls.Find = append(mock.calls.Find, callInfo)
	lockStatementViewStorageMockFind.Unlock()
	return mock.FindFunc(id)
}

// FindCalls gets all the calls that were made to Find.
// Check the length with:
//     len(mockedStatementViewStorage.FindCalls())
func (mock *StatementViewStorageMock) FindCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockStatementViewStorageMockFind.RLock()
	calls = mock.calls.Find
	lockStatementViewStorageMockFind.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *StatementViewStorageMock) Save(in1 ...model.EventRecord) error {
	if mock.SaveFunc == nil {
		panic("StatementViewStorageMock.SaveFunc: method is nil but StatementViewStorage.Save was just called")
	}
	callInfo := struct {
		In1 []model.EventRecord
	}{
		In1: in1,
	}
	lockStatementViewStorageMockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	lockStatementViewStorageMockSave.Unlock()
	return mock.SaveFunc(in1...)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//     len(mockedStatementViewStorage.SaveCalls())
func (mock *StatementViewStorageMock) SaveCalls() []struct {
	In1 []model.EventRecord
} {
	var calls []struct {
		In1 []model.EventRecord
	}
	lockStatementViewStorageMockSave.RLock()
	calls = mock.calls.Save
	lockStatementViewStorageMockSave.RUnlock()
	return calls
}

// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package usecases_test

import (
	"github.com/dc0d/workshop/model"
	"sync"
)

var (
	lockStatementViewRepositoryMockFind sync.RWMutex
)

// Ensure, that StatementViewRepositoryMock does implement model.StatementViewRepository.
// If this is not the case, regenerate this file with moq.
var _ model.StatementViewRepository = &StatementViewRepositoryMock{}

// StatementViewRepositoryMock is a mock implementation of model.StatementViewRepository.
//
//     func TestSomethingThatUsesStatementViewRepository(t *testing.T) {
//
//         // make and configure a mocked model.StatementViewRepository
//         mockedStatementViewRepository := &StatementViewRepositoryMock{
//             FindFunc: func(id string) (*model.Statement, error) {
// 	               panic("mock out the Find method")
//             },
//         }
//
//         // use mockedStatementViewRepository in code that requires model.StatementViewRepository
//         // and then make assertions.
//
//     }
type StatementViewRepositoryMock struct {
	// FindFunc mocks the Find method.
	FindFunc func(id string) (*model.Statement, error)

	// calls tracks calls to the methods.
	calls struct {
		// Find holds details about calls to the Find method.
		Find []struct {
			// ID is the id argument value.
			ID string
		}
	}
}

// Find calls FindFunc.
func (mock *StatementViewRepositoryMock) Find(id string) (*model.Statement, error) {
	if mock.FindFunc == nil {
		panic("StatementViewRepositoryMock.FindFunc: method is nil but StatementViewRepository.Find was just called")
	}
	callInfo := struct {
		ID string
	}{
		ID: id,
	}
	lockStatementViewRepositoryMockFind.Lock()
	mock.calls.Find = append(mock.calls.Find, callInfo)
	lockStatementViewRepositoryMockFind.Unlock()
	return mock.FindFunc(id)
}

// FindCalls gets all the calls that were made to Find.
// Check the length with:
//     len(mockedStatementViewRepository.FindCalls())
func (mock *StatementViewRepositoryMock) FindCalls() []struct {
	ID string
} {
	var calls []struct {
		ID string
	}
	lockStatementViewRepositoryMockFind.RLock()
	calls = mock.calls.Find
	lockStatementViewRepositoryMockFind.RUnlock()
	return calls
}

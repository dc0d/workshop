// Code generated by MockGen. DO NOT EDIT.
// Source: ./../domain_model/event_store.go

// Package repositories_test is a generated GoMock package.
package repositories_test

import (
	reflect "reflect"

	model "github.com/dc0d/workshop/domain_model"

	gomock "github.com/golang/mock/gomock"
)

// MockEventStore is a mock of EventStore interface
type MockEventStore struct {
	ctrl     *gomock.Controller
	recorder *MockEventStoreMockRecorder
}

// MockEventStoreMockRecorder is the mock recorder for MockEventStore
type MockEventStoreMockRecorder struct {
	mock *MockEventStore
}

// NewMockEventStore creates a new mock instance
func NewMockEventStore(ctrl *gomock.Controller) *MockEventStore {
	mock := &MockEventStore{ctrl: ctrl}
	mock.recorder = &MockEventStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEventStore) EXPECT() *MockEventStoreMockRecorder {
	return m.recorder
}

// Load mocks base method
func (m *MockEventStore) Load(streamID string) ([]model.EventRecord, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", streamID)
	ret0, _ := ret[0].([]model.EventRecord)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Load indicates an expected call of Load
func (mr *MockEventStoreMockRecorder) Load(streamID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockEventStore)(nil).Load), streamID)
}

// Append mocks base method
func (m *MockEventStore) Append(events ...model.EventRecord) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range events {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Append", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Append indicates an expected call of Append
func (mr *MockEventStoreMockRecorder) Append(events ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Append", reflect.TypeOf((*MockEventStore)(nil).Append), events...)
}

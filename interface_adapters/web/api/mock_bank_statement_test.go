// Code generated by MockGen. DO NOT EDIT.
// Source: ./../../../model/usecase_bank_statement.go

// Package api is a generated GoMock package.
package api

import (
	gomock "github.com/golang/mock/gomock"
	model "gitlab.com/dc0d/go-workshop/model"
	reflect "reflect"
)

// MockBankStatement is a mock of BankStatement interface
type MockBankStatement struct {
	ctrl     *gomock.Controller
	recorder *MockBankStatementMockRecorder
}

// MockBankStatementMockRecorder is the mock recorder for MockBankStatement
type MockBankStatementMockRecorder struct {
	mock *MockBankStatement
}

// NewMockBankStatement creates a new mock instance
func NewMockBankStatement(ctrl *gomock.Controller) *MockBankStatement {
	mock := &MockBankStatement{ctrl: ctrl}
	mock.recorder = &MockBankStatementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBankStatement) EXPECT() *MockBankStatementMockRecorder {
	return m.recorder
}

// Run mocks base method
func (m *MockBankStatement) Run(id string) (*model.Statement, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", id)
	ret0, _ := ret[0].(*model.Statement)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Run indicates an expected call of Run
func (mr *MockBankStatementMockRecorder) Run(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockBankStatement)(nil).Run), id)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: sqlhandler.go

// Package mock_database is a generated GoMock package.
package mock_database

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	database "github.com/ras0q/clean-architecture-sample/2_interface/database"
)

// MockSQLHandler is a mock of SQLHandler interface.
type MockSQLHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSQLHandlerMockRecorder
}

// MockSQLHandlerMockRecorder is the mock recorder for MockSQLHandler.
type MockSQLHandlerMockRecorder struct {
	mock *MockSQLHandler
}

// NewMockSQLHandler creates a new mock instance.
func NewMockSQLHandler(ctrl *gomock.Controller) *MockSQLHandler {
	mock := &MockSQLHandler{ctrl: ctrl}
	mock.recorder = &MockSQLHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSQLHandler) EXPECT() *MockSQLHandlerMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockSQLHandler) Create(value interface{}) database.SQLHandler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", value)
	ret0, _ := ret[0].(database.SQLHandler)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockSQLHandlerMockRecorder) Create(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockSQLHandler)(nil).Create), value)
}

// Error mocks base method.
func (m *MockSQLHandler) Error() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Error")
	ret0, _ := ret[0].(error)
	return ret0
}

// Error indicates an expected call of Error.
func (mr *MockSQLHandlerMockRecorder) Error() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockSQLHandler)(nil).Error))
}

// Find mocks base method.
func (m *MockSQLHandler) Find(out interface{}, where ...interface{}) database.SQLHandler {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Find", varargs...)
	ret0, _ := ret[0].(database.SQLHandler)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockSQLHandlerMockRecorder) Find(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockSQLHandler)(nil).Find), varargs...)
}

// First mocks base method.
func (m *MockSQLHandler) First(out interface{}, where ...interface{}) database.SQLHandler {
	m.ctrl.T.Helper()
	varargs := []interface{}{out}
	for _, a := range where {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "First", varargs...)
	ret0, _ := ret[0].(database.SQLHandler)
	return ret0
}

// First indicates an expected call of First.
func (mr *MockSQLHandlerMockRecorder) First(out interface{}, where ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{out}, where...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "First", reflect.TypeOf((*MockSQLHandler)(nil).First), varargs...)
}

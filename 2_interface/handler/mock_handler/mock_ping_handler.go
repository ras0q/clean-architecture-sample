// Code generated by MockGen. DO NOT EDIT.
// Source: ping_handler.go

// Package mock_handler is a generated GoMock package.
package mock_handler

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	handler "github.com/ras0q/clean-architecture-sample/2_interface/handler"
)

// MockPingHandler is a mock of PingHandler interface.
type MockPingHandler struct {
	ctrl     *gomock.Controller
	recorder *MockPingHandlerMockRecorder
}

// MockPingHandlerMockRecorder is the mock recorder for MockPingHandler.
type MockPingHandlerMockRecorder struct {
	mock *MockPingHandler
}

// NewMockPingHandler creates a new mock instance.
func NewMockPingHandler(ctrl *gomock.Controller) *MockPingHandler {
	mock := &MockPingHandler{ctrl: ctrl}
	mock.recorder = &MockPingHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPingHandler) EXPECT() *MockPingHandlerMockRecorder {
	return m.recorder
}

// Ping mocks base method.
func (m *MockPingHandler) Ping(c handler.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ping", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Ping indicates an expected call of Ping.
func (mr *MockPingHandlerMockRecorder) Ping(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*MockPingHandler)(nil).Ping), c)
}

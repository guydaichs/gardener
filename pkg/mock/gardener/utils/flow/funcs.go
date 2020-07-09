// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/mock/gardener/utils/flow (interfaces: TaskFn)

// Package flow is a generated GoMock package.
package flow

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTaskFn is a mock of TaskFn interface.
type MockTaskFn struct {
	ctrl     *gomock.Controller
	recorder *MockTaskFnMockRecorder
}

// MockTaskFnMockRecorder is the mock recorder for MockTaskFn.
type MockTaskFnMockRecorder struct {
	mock *MockTaskFn
}

// NewMockTaskFn creates a new mock instance.
func NewMockTaskFn(ctrl *gomock.Controller) *MockTaskFn {
	mock := &MockTaskFn{ctrl: ctrl}
	mock.recorder = &MockTaskFnMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTaskFn) EXPECT() *MockTaskFnMockRecorder {
	return m.recorder
}

// Do mocks base method.
func (m *MockTaskFn) Do(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Do indicates an expected call of Do.
func (mr *MockTaskFnMockRecorder) Do(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockTaskFn)(nil).Do), arg0)
}

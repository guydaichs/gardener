// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/extensions/pkg/gardener/chartrenderer (interfaces: Factory)

// Package chartrenderer is a generated GoMock package.
package chartrenderer

import (
	chartrenderer "github.com/gardener/gardener/pkg/chartrenderer"
	gomock "github.com/golang/mock/gomock"
	rest "k8s.io/client-go/rest"
	reflect "reflect"
)

// MockFactory is a mock of Factory interface.
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory.
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance.
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// NewForConfig mocks base method.
func (m *MockFactory) NewForConfig(arg0 *rest.Config) (chartrenderer.Interface, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewForConfig", arg0)
	ret0, _ := ret[0].(chartrenderer.Interface)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewForConfig indicates an expected call of NewForConfig.
func (mr *MockFactoryMockRecorder) NewForConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewForConfig", reflect.TypeOf((*MockFactory)(nil).NewForConfig), arg0)
}

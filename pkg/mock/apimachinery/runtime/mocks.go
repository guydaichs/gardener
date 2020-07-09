// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/apimachinery/pkg/runtime (interfaces: Decoder)

// Package runtime is a generated GoMock package.
package runtime

import (
	gomock "github.com/golang/mock/gomock"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	reflect "reflect"
)

// MockDecoder is a mock of Decoder interface.
type MockDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockDecoderMockRecorder
}

// MockDecoderMockRecorder is the mock recorder for MockDecoder.
type MockDecoderMockRecorder struct {
	mock *MockDecoder
}

// NewMockDecoder creates a new mock instance.
func NewMockDecoder(ctrl *gomock.Controller) *MockDecoder {
	mock := &MockDecoder{ctrl: ctrl}
	mock.recorder = &MockDecoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDecoder) EXPECT() *MockDecoderMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockDecoder) Decode(arg0 []byte, arg1 *schema.GroupVersionKind, arg2 runtime.Object) (runtime.Object, *schema.GroupVersionKind, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0, arg1, arg2)
	ret0, _ := ret[0].(runtime.Object)
	ret1, _ := ret[1].(*schema.GroupVersionKind)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Decode indicates an expected call of Decode.
func (mr *MockDecoderMockRecorder) Decode(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockDecoder)(nil).Decode), arg0, arg1, arg2)
}

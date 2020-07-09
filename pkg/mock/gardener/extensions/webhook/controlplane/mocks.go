// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/extensions/pkg/webhook/controlplane (interfaces: KubeletConfigCodec,UnitSerializer,FileContentInlineCodec)

// Package controlplane is a generated GoMock package.
package controlplane

import (
	unit "github.com/coreos/go-systemd/unit"
	v1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/kubelet/config/v1beta1"
	reflect "reflect"
)

// MockKubeletConfigCodec is a mock of KubeletConfigCodec interface.
type MockKubeletConfigCodec struct {
	ctrl     *gomock.Controller
	recorder *MockKubeletConfigCodecMockRecorder
}

// MockKubeletConfigCodecMockRecorder is the mock recorder for MockKubeletConfigCodec.
type MockKubeletConfigCodecMockRecorder struct {
	mock *MockKubeletConfigCodec
}

// NewMockKubeletConfigCodec creates a new mock instance.
func NewMockKubeletConfigCodec(ctrl *gomock.Controller) *MockKubeletConfigCodec {
	mock := &MockKubeletConfigCodec{ctrl: ctrl}
	mock.recorder = &MockKubeletConfigCodecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKubeletConfigCodec) EXPECT() *MockKubeletConfigCodecMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockKubeletConfigCodec) Decode(arg0 *v1alpha1.FileContentInline) (*v1beta1.KubeletConfiguration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0)
	ret0, _ := ret[0].(*v1beta1.KubeletConfiguration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockKubeletConfigCodecMockRecorder) Decode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockKubeletConfigCodec)(nil).Decode), arg0)
}

// Encode mocks base method.
func (m *MockKubeletConfigCodec) Encode(arg0 *v1beta1.KubeletConfiguration, arg1 string) (*v1alpha1.FileContentInline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.FileContentInline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode.
func (mr *MockKubeletConfigCodecMockRecorder) Encode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockKubeletConfigCodec)(nil).Encode), arg0, arg1)
}

// MockUnitSerializer is a mock of UnitSerializer interface.
type MockUnitSerializer struct {
	ctrl     *gomock.Controller
	recorder *MockUnitSerializerMockRecorder
}

// MockUnitSerializerMockRecorder is the mock recorder for MockUnitSerializer.
type MockUnitSerializerMockRecorder struct {
	mock *MockUnitSerializer
}

// NewMockUnitSerializer creates a new mock instance.
func NewMockUnitSerializer(ctrl *gomock.Controller) *MockUnitSerializer {
	mock := &MockUnitSerializer{ctrl: ctrl}
	mock.recorder = &MockUnitSerializerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnitSerializer) EXPECT() *MockUnitSerializerMockRecorder {
	return m.recorder
}

// Deserialize mocks base method.
func (m *MockUnitSerializer) Deserialize(arg0 string) ([]*unit.UnitOption, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deserialize", arg0)
	ret0, _ := ret[0].([]*unit.UnitOption)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deserialize indicates an expected call of Deserialize.
func (mr *MockUnitSerializerMockRecorder) Deserialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deserialize", reflect.TypeOf((*MockUnitSerializer)(nil).Deserialize), arg0)
}

// Serialize mocks base method.
func (m *MockUnitSerializer) Serialize(arg0 []*unit.UnitOption) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Serialize", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Serialize indicates an expected call of Serialize.
func (mr *MockUnitSerializerMockRecorder) Serialize(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Serialize", reflect.TypeOf((*MockUnitSerializer)(nil).Serialize), arg0)
}

// MockFileContentInlineCodec is a mock of FileContentInlineCodec interface.
type MockFileContentInlineCodec struct {
	ctrl     *gomock.Controller
	recorder *MockFileContentInlineCodecMockRecorder
}

// MockFileContentInlineCodecMockRecorder is the mock recorder for MockFileContentInlineCodec.
type MockFileContentInlineCodecMockRecorder struct {
	mock *MockFileContentInlineCodec
}

// NewMockFileContentInlineCodec creates a new mock instance.
func NewMockFileContentInlineCodec(ctrl *gomock.Controller) *MockFileContentInlineCodec {
	mock := &MockFileContentInlineCodec{ctrl: ctrl}
	mock.recorder = &MockFileContentInlineCodecMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileContentInlineCodec) EXPECT() *MockFileContentInlineCodecMockRecorder {
	return m.recorder
}

// Decode mocks base method.
func (m *MockFileContentInlineCodec) Decode(arg0 *v1alpha1.FileContentInline) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Decode", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Decode indicates an expected call of Decode.
func (mr *MockFileContentInlineCodecMockRecorder) Decode(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decode", reflect.TypeOf((*MockFileContentInlineCodec)(nil).Decode), arg0)
}

// Encode mocks base method.
func (m *MockFileContentInlineCodec) Encode(arg0 []byte, arg1 string) (*v1alpha1.FileContentInline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Encode", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.FileContentInline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Encode indicates an expected call of Encode.
func (mr *MockFileContentInlineCodecMockRecorder) Encode(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Encode", reflect.TypeOf((*MockFileContentInlineCodec)(nil).Encode), arg0, arg1)
}

// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package manager is a generated GoMock package.
package manager

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCryptFS is a mock of CryptFS interface.
type MockCryptFS struct {
	ctrl     *gomock.Controller
	recorder *MockCryptFSMockRecorder
}

// MockCryptFSMockRecorder is the mock recorder for MockCryptFS.
type MockCryptFSMockRecorder struct {
	mock *MockCryptFS
}

// NewMockCryptFS creates a new mock instance.
func NewMockCryptFS(ctrl *gomock.Controller) *MockCryptFS {
	mock := &MockCryptFS{ctrl: ctrl}
	mock.recorder = &MockCryptFSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCryptFS) EXPECT() *MockCryptFSMockRecorder {
	return m.recorder
}

// Open mocks base method.
func (m *MockCryptFS) Open() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockCryptFSMockRecorder) Open() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockCryptFS)(nil).Open))
}

// Write mocks base method.
func (m *MockCryptFS) Write(b []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", b)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockCryptFSMockRecorder) Write(b interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockCryptFS)(nil).Write), b)
}
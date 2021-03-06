// Code generated by MockGen. DO NOT EDIT.
// Source: access.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	iterator "github.com/syndtr/goleveldb/leveldb/iterator"
	util "github.com/syndtr/goleveldb/leveldb/util"
	reflect "reflect"
)

// MockAccess is a mock of Access interface
type MockAccess struct {
	ctrl     *gomock.Controller
	recorder *MockAccessMockRecorder
}

// MockAccessMockRecorder is the mock recorder for MockAccess
type MockAccessMockRecorder struct {
	mock *MockAccess
}

// NewMockAccess creates a new mock instance
func NewMockAccess(ctrl *gomock.Controller) *MockAccess {
	mock := &MockAccess{ctrl: ctrl}
	mock.recorder = &MockAccessMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccess) EXPECT() *MockAccessMockRecorder {
	return m.recorder
}

// Abort mocks base method
func (m *MockAccess) Abort() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Abort")
}

// Abort indicates an expected call of Abort
func (mr *MockAccessMockRecorder) Abort() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Abort", reflect.TypeOf((*MockAccess)(nil).Abort))
}

// Begin mocks base method
func (m *MockAccess) Begin() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Begin")
	ret0, _ := ret[0].(error)
	return ret0
}

// Begin indicates an expected call of Begin
func (mr *MockAccessMockRecorder) Begin() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Begin", reflect.TypeOf((*MockAccess)(nil).Begin))
}

// Commit mocks base method
func (m *MockAccess) Commit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockAccessMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockAccess)(nil).Commit))
}

// Delete mocks base method
func (m *MockAccess) Delete(arg0 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", arg0)
}

// Delete indicates an expected call of Delete
func (mr *MockAccessMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAccess)(nil).Delete), arg0)
}

// DumpTx mocks base method
func (m *MockAccess) DumpTx() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DumpTx")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// DumpTx indicates an expected call of DumpTx
func (mr *MockAccessMockRecorder) DumpTx() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DumpTx", reflect.TypeOf((*MockAccess)(nil).DumpTx))
}

// Get mocks base method
func (m *MockAccess) Get(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockAccessMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockAccess)(nil).Get), arg0)
}

// Has mocks base method
func (m *MockAccess) Has(arg0 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has
func (mr *MockAccessMockRecorder) Has(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockAccess)(nil).Has), arg0)
}

// InUse mocks base method
func (m *MockAccess) InUse() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InUse")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InUse indicates an expected call of InUse
func (mr *MockAccessMockRecorder) InUse() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InUse", reflect.TypeOf((*MockAccess)(nil).InUse))
}

// Iterator mocks base method
func (m *MockAccess) Iterator(arg0 *util.Range) iterator.Iterator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator", arg0)
	ret0, _ := ret[0].(iterator.Iterator)
	return ret0
}

// Iterator indicates an expected call of Iterator
func (mr *MockAccessMockRecorder) Iterator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockAccess)(nil).Iterator), arg0)
}

// Put mocks base method
func (m *MockAccess) Put(arg0, arg1 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Put", arg0, arg1)
}

// Put indicates an expected call of Put
func (mr *MockAccessMockRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*MockAccess)(nil).Put), arg0, arg1)
}

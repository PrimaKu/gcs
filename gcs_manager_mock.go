// Code generated by MockGen. DO NOT EDIT.
// Source: gcs_manager.go

// Package gcs is a generated GoMock package.
package gcs

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGCSManagerInterface is a mock of GCSManagerInterface interface.
type MockGCSManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockGCSManagerInterfaceMockRecorder
}

// MockGCSManagerInterfaceMockRecorder is the mock recorder for MockGCSManagerInterface.
type MockGCSManagerInterfaceMockRecorder struct {
	mock *MockGCSManagerInterface
}

// NewMockGCSManagerInterface creates a new mock instance.
func NewMockGCSManagerInterface(ctrl *gomock.Controller) *MockGCSManagerInterface {
	mock := &MockGCSManagerInterface{ctrl: ctrl}
	mock.recorder = &MockGCSManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGCSManagerInterface) EXPECT() *MockGCSManagerInterfaceMockRecorder {
	return m.recorder
}

// DeleteAllFilesInDirectory mocks base method.
func (m *MockGCSManagerInterface) DeleteAllFilesInDirectory(bucketName, directory string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllFilesInDirectory", bucketName, directory)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllFilesInDirectory indicates an expected call of DeleteAllFilesInDirectory.
func (mr *MockGCSManagerInterfaceMockRecorder) DeleteAllFilesInDirectory(bucketName, directory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllFilesInDirectory", reflect.TypeOf((*MockGCSManagerInterface)(nil).DeleteAllFilesInDirectory), bucketName, directory)
}

// DeleteFile mocks base method.
func (m *MockGCSManagerInterface) DeleteFile(bucketName, objectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", bucketName, objectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockGCSManagerInterfaceMockRecorder) DeleteFile(bucketName, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockGCSManagerInterface)(nil).DeleteFile), bucketName, objectName)
}

// ListFiles mocks base method.
func (m *MockGCSManagerInterface) ListFiles(bucketName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", bucketName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockGCSManagerInterfaceMockRecorder) ListFiles(bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockGCSManagerInterface)(nil).ListFiles), bucketName)
}

// MoveFile mocks base method.
func (m *MockGCSManagerInterface) MoveFile(bucketName, srcObjectName, dstObjectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveFile", bucketName, srcObjectName, dstObjectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveFile indicates an expected call of MoveFile.
func (mr *MockGCSManagerInterfaceMockRecorder) MoveFile(bucketName, srcObjectName, dstObjectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveFile", reflect.TypeOf((*MockGCSManagerInterface)(nil).MoveFile), bucketName, srcObjectName, dstObjectName)
}

// UploadFile mocks base method.
func (m *MockGCSManagerInterface) UploadFile(bucketName, objectName, filePath string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", bucketName, objectName, filePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockGCSManagerInterfaceMockRecorder) UploadFile(bucketName, objectName, filePath interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockGCSManagerInterface)(nil).UploadFile), bucketName, objectName, filePath)
}

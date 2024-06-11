// Code generated by MockGen. DO NOT EDIT.
// Source: gcs_manager.go

// Package gcs is a generated GoMock package.
package gcs

import (
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGCSManager is a mock of GCSManager interface.
type MockGCSManager struct {
	ctrl     *gomock.Controller
	recorder *MockGCSManagerMockRecorder
}

// MockGCSManagerMockRecorder is the mock recorder for MockGCSManager.
type MockGCSManagerMockRecorder struct {
	mock *MockGCSManager
}

// NewMockGCSManager creates a new mock instance.
func NewMockGCSManager(ctrl *gomock.Controller) *MockGCSManager {
	mock := &MockGCSManager{ctrl: ctrl}
	mock.recorder = &MockGCSManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGCSManager) EXPECT() *MockGCSManagerMockRecorder {
	return m.recorder
}

// DeleteAllFilesInDirectory mocks base method.
func (m *MockGCSManager) DeleteAllFilesInDirectory(bucketName, directory string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllFilesInDirectory", bucketName, directory)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllFilesInDirectory indicates an expected call of DeleteAllFilesInDirectory.
func (mr *MockGCSManagerMockRecorder) DeleteAllFilesInDirectory(bucketName, directory interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllFilesInDirectory", reflect.TypeOf((*MockGCSManager)(nil).DeleteAllFilesInDirectory), bucketName, directory)
}

// DeleteFile mocks base method.
func (m *MockGCSManager) DeleteFile(bucketName, objectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFile", bucketName, objectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteFile indicates an expected call of DeleteFile.
func (mr *MockGCSManagerMockRecorder) DeleteFile(bucketName, objectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFile", reflect.TypeOf((*MockGCSManager)(nil).DeleteFile), bucketName, objectName)
}

// ListFiles mocks base method.
func (m *MockGCSManager) ListFiles(bucketName string) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListFiles", bucketName)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListFiles indicates an expected call of ListFiles.
func (mr *MockGCSManagerMockRecorder) ListFiles(bucketName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListFiles", reflect.TypeOf((*MockGCSManager)(nil).ListFiles), bucketName)
}

// MoveFile mocks base method.
func (m *MockGCSManager) MoveFile(bucketName, srcObjectName, dstObjectName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MoveFile", bucketName, srcObjectName, dstObjectName)
	ret0, _ := ret[0].(error)
	return ret0
}

// MoveFile indicates an expected call of MoveFile.
func (mr *MockGCSManagerMockRecorder) MoveFile(bucketName, srcObjectName, dstObjectName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MoveFile", reflect.TypeOf((*MockGCSManager)(nil).MoveFile), bucketName, srcObjectName, dstObjectName)
}

// UploadFile mocks base method.
func (m *MockGCSManager) UploadFile(bucketName, objectName string, file os.File) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadFile", bucketName, objectName, file)
	ret0, _ := ret[0].(error)
	return ret0
}

// UploadFile indicates an expected call of UploadFile.
func (mr *MockGCSManagerMockRecorder) UploadFile(bucketName, objectName, file interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadFile", reflect.TypeOf((*MockGCSManager)(nil).UploadFile), bucketName, objectName, file)
}

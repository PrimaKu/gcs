package gcs

import (
	"io"
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUploadFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManager(ctrl)
	fileContent := []byte("test content")
	file := os.NewFile(0, "test-file")
	defer file.Close()
	file.Write(fileContent)
	file.Seek(0, io.SeekStart)

	mockGCS.EXPECT().UploadFile("bucketName", "objectName", *file).Return(nil)

	err := mockGCS.UploadFile("bucketName", "objectName", *file)
	assert.NoError(t, err)
}

func TestDeleteFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManager(ctrl)
	mockGCS.EXPECT().DeleteFile("bucketName", "objectName").Return(nil)

	err := mockGCS.DeleteFile("bucketName", "objectName")
	assert.NoError(t, err)
}

func TestListFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManager(ctrl)
	expectedFiles := []string{"file1", "file2"}
	mockGCS.EXPECT().ListFiles("bucketName", "/").Return(expectedFiles, nil)

	files, err := mockGCS.ListFiles("bucketName", "/")
	assert.NoError(t, err)
	assert.Equal(t, expectedFiles, files)
}

func TestMoveFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManager(ctrl)
	mockGCS.EXPECT().MoveFile("bucketName", "srcObjectName", "dstObjectName").Return(nil)

	err := mockGCS.MoveFile("bucketName", "srcObjectName", "dstObjectName")
	assert.NoError(t, err)
}

func TestDeleteAllFilesInDirectory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManager(ctrl)
	mockGCS.EXPECT().DeleteAllFilesInDirectory("bucketName", "directory/").Return(nil)

	err := mockGCS.DeleteAllFilesInDirectory("bucketName", "directory/")
	assert.NoError(t, err)
}

func TestDownloadFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManager(ctrl)
	mockGCS.EXPECT().DownloadFile("bucketName", "directory/file.ext", "/tmp/").Return(nil)

	err := mockGCS.DownloadFile("bucketName", "directory/file.ext", "/tmp/")
	assert.NoError(t, err)
}

func TestUploadDirectory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManager(ctrl)
	fileContent := []byte("test content")
	file := os.NewFile(0, "test-file")
	defer file.Close()
	file.Write(fileContent)
	file.Seek(0, io.SeekStart)

	mockGCS.EXPECT().UploadDirectory("bucketName", "localDir", "prefix").Return(nil)

	err := mockGCS.UploadDirectory("bucketName", "localDir", "prefix")
	assert.NoError(t, err)
}

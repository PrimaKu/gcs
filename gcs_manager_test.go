package gcs

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUploadFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManagerInterface(ctrl)
	mockGCS.EXPECT().UploadFile("bucketName", "objectName", "filePath").Return(nil)

	err := mockGCS.UploadFile("bucketName", "objectName", "filePath")
	assert.NoError(t, err)
}

func TestDeleteFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManagerInterface(ctrl)
	mockGCS.EXPECT().DeleteFile("bucketName", "objectName").Return(nil)

	err := mockGCS.DeleteFile("bucketName", "objectName")
	assert.NoError(t, err)
}

func TestListFiles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManagerInterface(ctrl)
	expectedFiles := []string{"file1", "file2"}
	mockGCS.EXPECT().ListFiles("bucketName").Return(expectedFiles, nil)

	files, err := mockGCS.ListFiles("bucketName")
	assert.NoError(t, err)
	assert.Equal(t, expectedFiles, files)
}

func TestMoveFile(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockGCS := NewMockGCSManagerInterface(ctrl)
	mockGCS.EXPECT().MoveFile("bucketName", "srcObjectName", "dstObjectName").Return(nil)

	err := mockGCS.MoveFile("bucketName", "srcObjectName", "dstObjectName")
	assert.NoError(t, err)
}

package gcs

import (
	"context"
	"fmt"
	"io"
	"os"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type (
	GCSManagerInterface interface {
		UploadFile(bucketName, objectName, filePath string) error
		DeleteFile(bucketName, objectName string) error
		ListFiles(bucketName string) ([]string, error)
	}

	GCSManager struct {
		client *storage.Client
		ctx    context.Context
	}
)

func NewGCSManager(credentialsFile string) (*GCSManager, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %v", err)
	}

	return &GCSManager{
		client: client,
		ctx:    ctx,
	}, nil
}

func (g *GCSManager) UploadFile(bucketName, objectName string, file os.File) error {
	bucket := g.client.Bucket(bucketName)
	object := bucket.Object(objectName)
	writer := object.NewWriter(g.ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, &file); err != nil {
		return fmt.Errorf("failed to upload file to GCS: %v", err)
	}

	return nil
}

func (g *GCSManager) DeleteFile(bucketName, objectName string) error {
	bucket := g.client.Bucket(bucketName)
	object := bucket.Object(objectName)
	if err := object.Delete(g.ctx); err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}

	return nil
}

func (g *GCSManager) ListFiles(bucketName string) ([]string, error) {
	var files []string
	it := g.client.Bucket(bucketName).Objects(g.ctx, nil)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to list objects: %v", err)
		}
		files = append(files, attrs.Name)
	}
	return files, nil
}

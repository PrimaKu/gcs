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
	GCSManager interface {
		UploadFile(bucketName, objectName string, file os.File) error
		DeleteFile(bucketName, objectName string) error
		ListFiles(bucketName string) ([]string, error)
		MoveFile(bucketName, srcObjectName, dstObjectName string) error
		DeleteAllFilesInDirectory(bucketName, directory string) error
	}

	gcsManager struct {
		client *storage.Client
		ctx    context.Context
	}
)

func NewGCSManager(credentialsFile string) (GCSManager, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile(credentialsFile))
	if err != nil {
		return nil, fmt.Errorf("failed to create GCS client: %v", err)
	}

	return &gcsManager{
		client: client,
		ctx:    ctx,
	}, nil
}

func (g gcsManager) UploadFile(bucketName, objectName string, file os.File) error {
	bucket := g.client.Bucket(bucketName)
	object := bucket.Object(objectName)
	writer := object.NewWriter(g.ctx)
	defer writer.Close()

	if _, err := io.Copy(writer, &file); err != nil {
		return fmt.Errorf("failed to upload file to GCS: %v", err)
	}

	return nil
}

func (g gcsManager) DeleteFile(bucketName, objectName string) error {
	bucket := g.client.Bucket(bucketName)
	object := bucket.Object(objectName)
	if err := object.Delete(g.ctx); err != nil {
		return fmt.Errorf("failed to delete object: %v", err)
	}

	return nil
}

func (g gcsManager) ListFiles(bucketName string) ([]string, error) {
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

func (g gcsManager) MoveFile(bucketName, srcObjectName, dstObjectName string) error {
	src := g.client.Bucket(bucketName).Object(srcObjectName)
	dst := g.client.Bucket(bucketName).Object(dstObjectName)

	_, err := dst.CopierFrom(src).Run(g.ctx)
	if err != nil {
		return err
	}

	if err := src.Delete(g.ctx); err != nil {
		return err
	}

	return nil
}

func (g gcsManager) DeleteAllFilesInDirectory(bucketName, directory string) error {
	bucket := g.client.Bucket(bucketName)
	query := &storage.Query{Prefix: directory}
	it := bucket.Objects(g.ctx, query)

	for {
		attrs, err := it.Next()
		if err == storage.ErrObjectNotExist {
			break
		}
		if err != nil {
			return err
		}

		obj := bucket.Object(attrs.Name)
		if err := obj.Delete(g.ctx); err != nil {
			return err
		}
	}

	return nil
}

package gcs

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type (
	GCSManager interface {
		UploadFile(bucketName, objectName string, file os.File) error
		DeleteFile(bucketName, objectName string) error
		ListFiles(bucketName, directory string) ([]string, error)
		MoveFile(bucketName, srcObjectName, dstObjectName string) error
		DeleteAllFilesInDirectory(bucketName, directory string) error
		DownloadFile(bucketName, objectName, destPath string) error
		UploadDirectory(bucketName, localDir, gcsPrefix string) error
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

func (g gcsManager) ListFiles(bucketName, directory string) ([]string, error) {
	if directory != "" && directory[len(directory)-1] != '/' {
		directory += "/"
	}

	var files []string
	it := g.client.Bucket(bucketName).Objects(g.ctx, &storage.Query{
		Prefix: directory,
	})
	for {
		attrs, err := it.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}

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
		if err != nil {
			if err == storage.ErrObjectNotExist {
				break
			}

			return err
		}

		obj := bucket.Object(attrs.Name)
		if err := obj.Delete(g.ctx); err != nil {
			return err
		}
	}

	return nil
}

func (g gcsManager) DownloadFile(bucketName, objectName, destPath string) error {
	bucket := g.client.Bucket(bucketName)
	obj := bucket.Object(objectName)
	r, err := obj.NewReader(g.ctx)
	if err != nil {
		return err
	}
	defer r.Close()

	out, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, r)
	return err
}

func (g gcsManager) UploadDirectory(bucketName, localDir, gcsPrefix string) error {
	err := filepath.Walk(localDir, func(localPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		localPath = strings.ReplaceAll(localPath, "\\", "/")
		localDir = strings.ReplaceAll(localDir, "\\", "/")

		relativePath := strings.TrimPrefix(localPath, strings.Replace(localDir, "./", "", 1))
		relativePath = strings.TrimPrefix(relativePath, string(filepath.Separator))
		gcsPath := filepath.Join(gcsPrefix, relativePath)

		file, err := os.Open(localPath)
		if err != nil {
			return fmt.Errorf("os.Open %v: %v", localPath, err)
		}
		defer file.Close()

		writer := g.client.Bucket(bucketName).Object(gcsPath).NewWriter(g.ctx)
		if _, err := io.Copy(writer, file); err != nil {
			return fmt.Errorf("io.Copy: %v", err)
		}

		if err := writer.Close(); err != nil {
			return fmt.Errorf("writer.Close: %v", err)
		}

		fmt.Printf("Uploaded %s to gs://%s/%s\n", localPath, bucketName, gcsPath)
		return nil
	})

	if err != nil {
		return fmt.Errorf("filepath.Walk: %v", err)
	}

	return nil
}

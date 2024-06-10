# GCS Manager

A Go library for managing Google Cloud Storage (GCS).

## Installation

```bash
go get github.com/PrimaKu/gcs
```

## Setup

```go
gcsManager, err := gcs.NewGCSManager(credentialPath) // string
if err != nil {
  log.Fatalf("Failed to create GCS client: %v", err)
}
```

## Upload File
```go
gcsManager.UploadFile(bucketName, fileName, file) // string, string, os.File
```
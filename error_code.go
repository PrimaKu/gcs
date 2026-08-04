package gcs

import "cloud.google.com/go/storage"

var ErrFileNotExist = storage.ErrObjectNotExist
var ErrBucketNotExist = storage.ErrBucketNotExist

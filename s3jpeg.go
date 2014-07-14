package s3jpeg

import (
  "github.com/mitchellh/goamz/aws"
  "github.com/mitchellh/goamz/s3"
)

type Bucket s3.Bucket

func New(name string) *Bucket {
  bucket := new(Bucket)
  bucket.Name = name

  auth, e := aws.EnvAuth()
  if e == nil {
    s3 := new(s3.S3)
    s3.Auth = auth
    bucket.S3 = s3
  }

  return bucket
}

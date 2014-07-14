package s3jpeg

import (
  "image"
  "image/jpeg"
  "github.com/mitchellh/goamz/aws"
  "github.com/mitchellh/goamz/s3"
)

type Bucket struct {
  s3.Bucket
}

func New(name string) *Bucket {
  bucket := new(Bucket)
  bucket.Name = name

  auth, e := aws.EnvAuth()
  if e == nil {
    s3 := new(s3.S3)
    s3.Auth = auth
    s3.Region = aws.USEast
    bucket.S3 = s3
  }

  return bucket
}

func (b *Bucket) GetImage(path string) (image.Image, error) {
  reader, e := b.GetReader(path)
  if e != nil {
    return nil, e
  }
  return jpeg.Decode(reader)
}

package s3jpeg

import (
  "image"
  "image/jpeg"
  "bytes"
  "github.com/mitchellh/goamz/aws"
  "github.com/mitchellh/goamz/s3"
  "github.com/disintegration/imaging"
)

type Bucket struct {
  s3.Bucket
  ACL s3.ACL
}

type Version struct {
  Width, Height, Quality int
}

func New(name string) *Bucket {
  bucket := new(Bucket)
  bucket.Name = name
  bucket.ACL = s3.PublicRead

  
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

func (b *Bucket) PutImage(path string, i image.Image, quality int) error {
  var r bytes.Buffer
  e := jpeg.Encode(&r, i, &jpeg.Options{ quality })
  if e != nil {
    return e
  }
  len := int64(r.Len())
  return b.PutReader(path, &r, len, "image/jpeg", b.ACL)
}

func (b *Bucket) PutThumbnail(path string, i image.Image, v *Version) error {
  img := imaging.Thumbnail(i, v.Width, v.Height, imaging.CatmullRom)
  return b.PutImage(path, img, v.Quality)
}

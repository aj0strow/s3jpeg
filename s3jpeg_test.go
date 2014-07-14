package s3jpeg

import (
  "testing"
  "os"
  "image"
  "image/jpeg"
)

func openImage(fname string) (image.Image, error) {
  file, e := os.Open(fname)
  if e != nil {
    return nil, e
  }
  return jpeg.Decode(file)
}

func TestNew(t *testing.T) {
  bucket := New("modde-test")
  if bucket.S3 == nil {
    t.Errorf("Coundn't find auth in env.")
  }
}

func TestGetImage(t *testing.T) {
  bucket := New("modde-test")
  _, e := bucket.GetImage("test_image.jpg")
  if e != nil {
    t.Errorf("Got error %s.", e.Error())
  }
}

func TestPutImage(t *testing.T) {
  bucket := New("modde-test")
  img, e := openImage("logo.jpg")
  e = bucket.PutImage("test/logo.jpg", img, 50)
  if e != nil {
    t.Errorf("Got error %s.", e.Error())
  }
}

func TestPutThumbnail(t *testing.T) {
  bucket := New("modde-test")
  img, e := openImage("logo.jpg")
  e = bucket.PutThumbnail("test/logo_small.jpg", img, &Version{ 100, 50, 80 })
  if e != nil {
    t.Errorf("Got error %s.", e.Error())
  }
}

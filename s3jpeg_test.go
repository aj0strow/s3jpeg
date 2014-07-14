package s3jpeg

import (
  "testing"
  "os"
  "image/jpeg"
)

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
  file, e := os.Open("logo.jpg")
  if e != nil {
    t.Errorf("Can't open modde.jpg file.")
  }
  img, e := jpeg.Decode(file)
  if e != nil {
    t.Errorf("Got decode error %s.", e.Error())
  }
  e = bucket.PutImage("test/logo.jpg", img, 50)
  if e != nil {
    t.Errorf("Got error %s.", e.Error())
  }
}

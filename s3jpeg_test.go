package s3jpeg

import (
  "testing"
)

func TestNew(t *testing.T) {
  bucket := New("modde-test")
  if bucket.S3 == nil {
    t.Errorf("Coundn't find auth in env.")
  }
}

func TestGetImage(t *testing.T) {
  bucket := New("modde-test")
  jpeg, e := bucket.GetImage("test_image.jpg")
  if jpeg == nil || e != nil {
    t.Errorf("Got error %s.", e.Error())
  }
}

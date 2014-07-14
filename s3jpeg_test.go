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

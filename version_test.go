package s3jpeg

import (
  "testing"
)

func TestNewVersion(t *testing.T) {
  v := NewVersion("40x60:80")
  if v.Width != 40 {
    t.Errorf("Width bad.")
  }
  if v.Height != 60 {
    t.Errorf("Height bad.")
  }
  if v.Quality != 80 {
    t.Errorf("Quality bad.")
  }
}

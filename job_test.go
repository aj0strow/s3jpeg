package s3jpeg

import (
	"testing"
)

var DEFAULT_QUALITY = 90

func TestNewJob(t *testing.T) {
  job := NewJob(Resize, "100x0:75")
  if job.Version.Width != 100 {
    t.Errorf("Width should be %d not %d.", 100, job.Version.Width)
  }
  if job.Version.Height != 0 {
    t.Errorf("Height should be %d not %d.", 0, job.Version.Height)
  }
  if job.Version.Quality != 75 {
    t.Errorf("Quality should be %d not %d.", 75, job.Version.Quality)
  }
}

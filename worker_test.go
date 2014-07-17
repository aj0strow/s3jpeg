package s3jpeg

import (
  "testing"
  "os"
)

var bucket = os.Getenv("AWS_S3_BUCKET")

func TestNewWorker(t *testing.T) {
  worker := NewWorker(bucket)
  if worker.Bucket == nil {
    t.Errorf("Bucket nil.")
  }
  if len(worker.Versions) != 0 {
    t.Errorf("Versions wrong length.")
  }
}

func TestQueue(t *testing.T) {
  worker := NewWorker(bucket)
  worker.Queue("50x50:75")
  if len(worker.Versions) != 1 {
    t.Errorf("Queue failed.")
  }
}

func TestRun(t *testing.T) {
  worker := NewWorker(bucket)
  r, e := worker.Bucket.GetReader("test/logo.jpg")
  if e != nil {
    t.Errorf(e.Error())
  }
  worker.Queue("test/logo_50x50:75.jpg")
  worker.Queue("test/logo_60x40:75.jpg")
  e = worker.Run(r)
  if e != nil {
    t.Errorf(e.Error())
  }
}

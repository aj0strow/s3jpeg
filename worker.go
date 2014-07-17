package s3jpeg

import (
  "io"
  "image"
  "image/jpeg"
)

type Worker struct {
  Bucket *Bucket
  Versions []*Version
}

func NewWorker(name string) *Worker {
  worker := new(Worker)
  worker.Bucket = NewBucket(name)
  return worker
}

func (w *Worker) Queue(v string) {
  w.Versions = append(w.Versions, NewVersion(v))
}

func (w *Worker) Run(r io.Reader) error {
	i, e := jpeg.Decode(r)
	if e != nil {
		return e
	}
	return w.RunImage(i)
}

func (w *Worker) RunImage(i image.Image) error {
  c := make(chan error)

  for _, v := range w.Versions {
    go w.create(i, v, c)
  }

  for _ = range w.Versions {
    if e := <- c; e != nil {
      return e
    }
  }
  return nil
}

func (w *Worker) create(i image.Image, v *Version, c chan error) {
  c <- w.Bucket.PutThumbnailImage(i, v)
}

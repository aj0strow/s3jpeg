package s3jpeg

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

package s3jpeg

import (
	"image"
	"image/jpeg"
	"os"
	"testing"
)

func openImage(fname string) (image.Image, error) {
	file, e := os.Open(fname)
	defer file.Close()
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

func TestPutResize(t *testing.T) {
	bucket := New("modde-test")
	img, e := openImage("logo.jpg")
	small := Version{100, 0, 80}
	e = bucket.PutResize("test/resize/image.jpg", img, &small)
	if e != nil {
		t.Errorf("Got error %s.", e.Error())
	}
}

func TestPutResizeReader(t *testing.T) {
	bucket := New("modde-test")
	file, e := os.Open("logo.jpg")
	medium := Version{100, 0, 80}
	e = bucket.PutResizeReader("test/resize/reader.jpg", file, &medium)
	if e != nil {
		t.Errorf("Got error %s.", e.Error())
	}
}

func TestPutFit(t *testing.T) {
	bucket := New("modde-test")
	img, e := openImage("logo.jpg")
	small := Version{200, 200, 60}
	e = bucket.PutFit("test/fit/image.jpg", img, &small)
	if e != nil {
		t.Errorf("Got error %s.", e.Error())
	}
}

func TestPutFitReader(t *testing.T) {
	bucket := New("modde-test")
	file, e := os.Open("logo.jpg")
	medium := Version{200, 200, 60}
	e = bucket.PutFitReader("test/fit/reader.jpg", file, &medium)
	if e != nil {
		t.Errorf("Got error %s.", e.Error())
	}
}

func TestPutThumbnail(t *testing.T) {
	bucket := New("modde-test")
	img, e := openImage("logo.jpg")
	small := Version{100, 45, 80}
	e = bucket.PutThumbnail("test/thumbnail/image.jpg", img, &small)
	if e != nil {
		t.Errorf("Got error %s.", e.Error())
	}
}

func TestPutThumbnailReader(t *testing.T) {
	bucket := New("modde-test")
	file, e := os.Open("logo.jpg")
	medium := Version{100, 45, 90}
	e = bucket.PutThumbnailReader("test/thumbnail/reader.jpg", file, &medium)
	if e != nil {
		t.Errorf("Got error %s.", e.Error())
	}
}

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

func TestNewBucket(t *testing.T) {
	bucket := NewBucket("modde-test")
	if bucket.S3 == nil {
		t.Errorf("Coundn't find auth in env.")
	}
}

func TestGetImage(t *testing.T) {
	bucket := NewBucket("modde-test")
	_, e := bucket.GetImage("test_image.jpg")
	if e != nil {
		t.Errorf("Got error %s.", e.Error())
	}
}

func TestPutImage(t *testing.T) {
	bucket := NewBucket("modde-test")
	img, e := openImage("logo.jpg")
	e = bucket.PutImage("test/logo.jpg", img, 50)
	if e != nil {
		t.Errorf("Got error %s.", e.Error())
	}
}

func TestPutThumbnailImage(t *testing.T) {
	bucket := NewBucket("modde-test")
	img, e := openImage("logo.jpg")
	small := Version{"test/thumbnail/image.jpg", 100, 45, 80}
	e = bucket.PutThumbnailImage(img, &small)
	if e != nil {
		t.Errorf("Got error %s.", e.Error())
	}
}

func TestPutThumbnail(t *testing.T) {
	bucket := NewBucket("modde-test")
	file, e := os.Open("logo.jpg")
	medium := Version{"test/thumbnail/reader.jpg", 100, 45, 90}
	e = bucket.PutThumbnail(file, &medium)
	if e != nil {
		t.Errorf("Got error %s.", e.Error())
	}
}

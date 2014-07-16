package s3jpeg

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
	"image"
	"image/jpeg"
	"io"
)

type Bucket struct {
	s3.Bucket
	ACL s3.ACL
}

func NewBucket(name string) *Bucket {
	bucket := new(Bucket)
	bucket.Name = name
	bucket.ACL = s3.PublicRead

	auth, e := aws.EnvAuth()
	if e == nil {
		s3 := new(s3.S3)
		s3.Auth = auth
		s3.Region = aws.USEast
		bucket.S3 = s3
	}

	return bucket
}

// GetImage & PutImage
// wrap s3.Bucket methods to deal with jpeg image files

func (b *Bucket) GetImage(path string) (image.Image, error) {
	reader, e := b.GetReader(path)
	if e != nil {
		return nil, e
	}
	return jpeg.Decode(reader)
}

func (b *Bucket) PutImage(path string, i image.Image, quality int) error {
	var r bytes.Buffer
	e := jpeg.Encode(&r, i, &jpeg.Options{quality})
	if e != nil {
		return e
	}
	len := int64(r.Len())
	return b.PutReader(path, &r, len, "image/jpeg", b.ACL)
}

// Resize
// leave width or height 0 to maintain ratio

func (b *Bucket) PutResize(path string, i image.Image, v *Version) error {
	img := imaging.Resize(i, v.Width, v.Height, imaging.Lanczos)
	return b.PutImage(path, img, v.Quality)
}

func (b *Bucket) PutResizeReader(path string, r io.Reader, v *Version) error {
	img, e := jpeg.Decode(r)
	if e != nil {
		return e
	}
	return b.PutResize(path, img, v)
}

// Fit
// scale image down to fit within width x height box

func (b *Bucket) PutFit(path string, i image.Image, v *Version) error {
	img := imaging.Fit(i, v.Width, v.Height, imaging.Lanczos)
	return b.PutImage(path, img, v.Quality)
}

func (b *Bucket) PutFitReader(path string, r io.Reader, v *Version) error {
	img, e := jpeg.Decode(r)
	if e != nil {
		return e
	}
	return b.PutFit(path, img, v)
}

// Thumbnail
// crop image to width and height

func (b *Bucket) PutThumbnail(path string, i image.Image, v *Version) error {
	img := imaging.Thumbnail(i, v.Width, v.Height, imaging.Lanczos)
	return b.PutImage(path, img, v.Quality)
}

func (b *Bucket) PutThumbnailReader(path string, r io.Reader, v *Version) error {
	img, e := jpeg.Decode(r)
	if e != nil {
		return e
	}
	return b.PutThumbnail(path, img, v)
}

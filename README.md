# s3jpeg

> Parallel S3 Thumbnails

This tool is opinionated. I use it for an [application with user uploaded pictures and frequently changing UI](http://modde.co/). Space is cheap so it makes sense to generate thumbnail versions to improve speed.

Different thumbnail versions are often named with suffixes such as `_large.jpg` or `_small.jpg`. I think it's better to explicitly state the size and lossy quality like `_100x100:80.jpg`.

### Example

Take an `io.Reader` then queue and run thumbnail jobs in parallel. 

```go
// example image id
id := "9283lkajwfeawef0000"

// somehow get an io reader
image, err := os.Open(fmt.Sprintf("tmp/%s.jpg", id))

// create worker
worker := s3jpeg.NewWorker("s3-bucket-name")

// queue parallel thumbnail jobs
worker.Queue(fmt.Sprintf("images/%s/600x600:75.jpg", id))
worker.Queue(fmt.Sprintf("images/%s/160x160:85.jpg", id))
worker.Queue(fmt.Sprintf("images/%s/160x160:45.jpg", id))
worker.Queue(fmt.Sprintf("images/%s/50x50:75.jpg", id))

// generate images and store to s3
err = worker.Run(image)

if err != nil {
  // handle error
}
```

It seems to take 1 or 2 seconds. I wrote this in a cafe tho so hopefully it's faster. 

### Caveats

There are a number of limitations because my use case is limited. Feel free to fork. 

* It only works on image files with the `width x height : quality` format.
* It only works with JPEGs.
* It only works with AWS S3 storage.
* It only works with environment variables for key & secret.

### Notes

Built on top of the excellent **[imaging](https://github.com/disintegration/imaging)** and **[goamz](https://github.com/mitchellh/goamz)** libraries. 

```
$ go get github.com/disintegration/imaging
$ go get github.com/mitchellh/goamz
$ go get github.com/aj0strow/s3jpeg
```

Next moves is the CLI interface. 

License: **MIT**

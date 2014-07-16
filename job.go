package s3jpeg

import (
  "regexp"
  "strings"
  "strconv"
)

type Strategy int

const (
  Thumbnail = 0
  Fit = 1
  Resize = 2
)

type Version struct {
	Width, Height, Quality int
}

type Job struct {
  Version *Version
  Key string
  Strategy Strategy
}

func NewJob(s Strategy, key string) *Job {
  job := new(Job)
  job.Version = parseVersion(key)
  if !strings.HasSuffix(key, ".jpg") {
    key += ".jpg"
  }
  job.Key = key
  job.Strategy = s
  return job
}

func parseVersion (key string) *Version {
  version := new(Version)

  re := regexp.MustCompile("(\\d+)x(\\d+):(\\d+)")
  match := re.FindAllStringSubmatch(key, -1)
  strs := match[0][1:]

  version.Width = parseInt(strs[0])
  version.Height = parseInt(strs[1])
  version.Quality = parseInt(strs[2])

  return version
}

func parseInt (s string) int {
  val, _ := strconv.ParseInt(s, 10, 64)
  return int(val)
}

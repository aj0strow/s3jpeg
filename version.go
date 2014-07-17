package s3jpeg

import (
  "regexp"
  "strings"
  "strconv"
)

type Version struct {
  Key string // relative s3 file key
  Width, Height, Quality int
}

func NewVersion(key string) *Version {
  version := new(Version)

  if !strings.HasSuffix(key, ".jpg") {
    key += ".jpg"
  }
  version.Key = key

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

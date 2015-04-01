package md5

import (
	"os"
)

type downloaderMd5 struct {
}

func New() *downloaderMd5 {
	d := new(downloaderMd5)
	return d
}

func (d *downloaderMd5) Download(url string, targetDirectory *os.File) error {
	return nil
}

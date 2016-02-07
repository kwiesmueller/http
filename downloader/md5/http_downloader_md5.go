package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"


	http_util "github.com/bborbe/http/util"
	io_file_writer "github.com/bborbe/io/file_writer"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type GetUrl func(url string) (resp *http.Response, err error)

type downloaderMd5 struct {
	getUrl GetUrl
}

func New(getUrl GetUrl) *downloaderMd5 {
	d := new(downloaderMd5)
	d.getUrl = getUrl
	return d
}

func (d *downloaderMd5) Download(url string, targetDirectory *os.File) error {
	return download(url, targetDirectory, d.getUrl)
}

func download(url string, targetDirectory *os.File, getUrl GetUrl) error {
	logger.Debugf("download %s to directory %s", url, targetDirectory.Name())
	response, err := getUrl(url)
	if err != nil {
		return err
	}
	content, err := http_util.ResponseToByteArray(response)
	if err != nil {
		return err
	}
	filename := createFilename(content, response, targetDirectory)
	logger.Debugf("filename: %s", filename)
	return saveToFile(content, filename)
}

func createFilename(content []byte, response *http.Response, directory *os.File) string {
	logger.Debugf("createFilename")
	md5string := createMd5Checksum(content)
	ext := http_util.FindFileExtension(response)
	return fmt.Sprintf("%s%c%s.%s", directory.Name(), os.PathSeparator, md5string, ext)
}

func createMd5Checksum(content []byte) string {
	logger.Debugf("create md5 checksum")
	hasher := md5.New()
	hasher.Write(content)
	return hex.EncodeToString(hasher.Sum(nil))
}

func saveToFile(content []byte, filename string) error {
	logger.Debugf("save content to %s", filename)
	writer, err := io_file_writer.NewFileWriter(filename)
	defer writer.Close()
	if err != nil {
		return err
	}
	writer.Write(content)
	return err
}

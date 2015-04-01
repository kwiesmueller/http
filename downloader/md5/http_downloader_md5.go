package md5

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"

	http_client "github.com/bborbe/http/client"
	http_util "github.com/bborbe/http/util"
	"github.com/bborbe/io/file_writer"
	"github.com/bborbe/log"
	"github.com/bborbe/stringutil"
)

var logger = log.DefaultLogger

type downloaderMd5 struct {
	getDownloader http_client.GetDownloader
}

func New(getDownloader http_client.GetDownloader) *downloaderMd5 {
	d := new(downloaderMd5)
	d.getDownloader = getDownloader
	return d
}

func (d *downloaderMd5) Download(url string, targetDirectory *os.File) error {
	return download(url, targetDirectory, d.getDownloader)
}

func download(url string, targetDirectory *os.File, getDownloader http_client.GetDownloader) error {
	response, err := getDownloader.Get(url)
	if err != nil {
		return err
	}
	content, err := http_util.ResponseToByteArray(response)
	if err != nil {
		return err
	}
	filename := createFilename(content, response, targetDirectory)
	logger.Tracef("filename: %s", filename)
	return saveToFile(content, filename)
}

func createFilename(content []byte, response *http.Response, directory *os.File) string {
	md5string := createMd5Checksum(content)
	ext := getExt(response)
	return fmt.Sprintf("%s%c%s.%s", directory.Name(), os.PathSeparator, md5string, ext)
}

func createMd5Checksum(content []byte) string {
	hasher := md5.New()
	hasher.Write(content)
	return hex.EncodeToString(hasher.Sum(nil))
}

func getExt(response *http.Response) string {
	path := response.Request.URL.Path
	return stringutil.StringAfter(path, ".")
}

func saveToFile(content []byte, filename string) error {
	writer, err := file_writer.NewFileWriter(filename)
	if err != nil {
		return err
	}
	writer.Write(content)
	writer.Close()
	return err
}

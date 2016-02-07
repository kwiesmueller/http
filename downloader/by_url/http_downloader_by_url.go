package by_url

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	io_file_writer "github.com/bborbe/io/file_writer"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type DownloadUrl func(url string) (resp *http.Response, err error)

type downloaderByUrl struct {
	downloadUrl DownloadUrl
}

func New(downloadUrl DownloadUrl) *downloaderByUrl {
	d := new(downloaderByUrl)
	d.downloadUrl = downloadUrl
	return d
}

func (d *downloaderByUrl) Download(url string, targetDirectory *os.File) error {
	return downloadLink(url, targetDirectory, d.downloadUrl)
}

func downloadLink(url string, targetDirectory *os.File, downloadUrl DownloadUrl) error {
	logger.Debugf("download %s started", url)
	response, err := downloadUrl(url)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		content, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		logger.Debugf("%s", string(content))
		return errors.New(string(content))
	}

	targetDirectory.Name()

	filename := createFilename(url)
	logger.Debugf("to %s", filename)
	writer, err := io_file_writer.NewFileWriter(fmt.Sprintf("%s/%s", targetDirectory.Name(), filename))
	if err != nil {
		logger.Errorf("open '%s' failed", filename)
		return err
	}
	io.Copy(writer, response.Body)
	writer.Flush()
	writer.Close()
	logger.Debugf("download %s finished", url)
	return nil
}

func createFilename(url string) string {
	re := regexp.MustCompile("[^A-Za-z0-9\\.]+")
	return re.ReplaceAllString(url, "_")
}

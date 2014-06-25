package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bborbe/io/file_writer"
	"github.com/bborbe/log"
	"regexp"
)

var logger = log.DefaultLogger

const SIZE = 10

func main() {
	logger.SetLevelThreshold(log.ERROR)
	defer logger.Close()
	logger.Debug("started")
	err := downloadLinks()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("finished")
}

func downloadLinks() error {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			return err
		}
		err = downloadLink(string(line))
		if err != nil {
			return err
		}
	}
	return nil
}

func downloadLink(url string) error {
	logger.Debugf("download %s started", url)
	response, err := http.Get(url)
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
	filename := createFilename(url)
	logger.Debugf("to %s", filename)
	writer, err := file_writer.NewFileWriter(filename)
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
	return re.ReplaceAllString(url,"_")
}

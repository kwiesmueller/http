package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"github.com/bborbe/log"
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
	io.Copy(os.Stdout, response.Body)
	logger.Debugf("download %s finished", url)
	return nil
}

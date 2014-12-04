package main

import (
	"bytes"
	"io"
	"os"

	"github.com/bborbe/crawler/linkparser"
	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

const SIZE = 10

func main() {
	logger.SetLevelThreshold(log.ERROR)
	defer logger.Close()
	logger.Debug("started")
	err := grepUrls()
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug("finished")
}

func grepUrls() error {
	contentBuffer := bytes.NewBuffer(nil)
	io.Copy(contentBuffer, os.Stdin)

	l := linkparser.New()
	links := l.Parse(string(contentBuffer.Bytes()))
	for match := range links {
		os.Stdout.WriteString(match)
		os.Stdout.WriteString("\n")
	}

	return nil
}

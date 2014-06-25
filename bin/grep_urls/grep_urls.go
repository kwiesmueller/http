package main

import (    "io"
	"os"
	"github.com/bborbe/log"
	"bytes"
	"regexp"
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
	re := regexp.MustCompile("(https?://[^,'\" \\?&;]+)")
	result := re.FindAll(contentBuffer.Bytes(), -1)
	for _, match := range result {
		os.Stdout.Write(match)
		os.Stdout.WriteString("\n")
	}
	return nil
}


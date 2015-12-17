package main

import (
	"bytes"
	"io"
	"os"

	"flag"

	crawler_linkparser "github.com/bborbe/crawler/linkparser"

	"fmt"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

const (
	PARAMETER_LOGLEVEL = "loglevel"
)

func main() {
	defer logger.Close()
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	writer := os.Stdout
	input := os.Stdin
	err := do(writer, input)
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, input io.Reader) error {
	contentBuffer := bytes.NewBuffer(nil)
	io.Copy(contentBuffer, input)

	linkparser := crawler_linkparser.New()
	links := linkparser.ParseAbsolute(string(contentBuffer.Bytes()))
	for match := range links {
		fmt.Fprintf(writer, "%s\n", match)
	}

	return nil
}

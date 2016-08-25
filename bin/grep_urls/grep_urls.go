package main

import (
	"bytes"
	"io"
	"os"

	"flag"

	crawler_linkparser "github.com/bborbe/crawler/linkparser"

	"fmt"

	"runtime"

	"github.com/bborbe/log"
)

const (
	PARAMETER_LOGLEVEL = "loglevel"
)

var (
	logger      = log.DefaultLogger
	logLevelPtr = flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
)

func main() {
	defer logger.Close()
	flag.Parse()

	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

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
	if _, err := io.Copy(contentBuffer, input); err != nil {
		return err
	}
	linkparser := crawler_linkparser.New()
	links := linkparser.ParseAbsolute(string(contentBuffer.Bytes()))
	for match := range links {
		fmt.Fprintf(writer, "%s\n", match)
	}

	return nil
}

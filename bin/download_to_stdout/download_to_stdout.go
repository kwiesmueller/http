package main

import (
	"bufio"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"flag"

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
	reader := bufio.NewReader(input)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = downloadLink(writer, string(line))
		if err != nil {
			return err
		}
	}
}

func downloadLink(writer io.Writer, url string) error {
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
	if _, err := io.Copy(writer, response.Body); err != nil {
		return err
	}
	logger.Debugf("download %s finished", url)
	return nil
}

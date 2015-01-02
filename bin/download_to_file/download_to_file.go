package main

import (
	"bufio"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"regexp"

	"flag"
	"io"

	"github.com/bborbe/io/file_writer"
	"github.com/bborbe/log"

	"runtime"
	"sync"
)

var logger = log.DefaultLogger

const (
	PARAMETER_LOGLEVEL           = "loglevel"
	PARAMETER_PARALLEL_DOWNLOADS = "max"
	DEFAULT_PARALLEL_DOWNLOADS   = 2
)

func main() {
	defer logger.Close()
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, "one of OFF,TRACE,DEBUG,INFO,WARN,ERROR")
	maxConcurrencyDownloadsPtr := flag.Int(PARAMETER_PARALLEL_DOWNLOADS, DEFAULT_PARALLEL_DOWNLOADS, "max parallel downloads")
	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

	writer := os.Stdout
	input := os.Stdin
	wg := new(sync.WaitGroup)
	err := do(writer, input, *maxConcurrencyDownloadsPtr, wg)
	wg.Wait()
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, input io.Reader, maxConcurrencyDownloads int, wg *sync.WaitGroup) error {
	throttle := make(chan bool, maxConcurrencyDownloads)
	reader := bufio.NewReader(input)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		wg.Add(1)
		go func() {
			link := string(line)
			throttle <- true
			downloadLink(link)
			<-throttle
			wg.Done()
		}()
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
	return re.ReplaceAllString(url, "_")
}

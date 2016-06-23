package main

import (
	"bufio"
	"os"

	"flag"
	"io"

	"github.com/bborbe/log"

	"fmt"
	"runtime"
	"sync"

	http_client_builder "github.com/bborbe/http/client_builder"
	http_downloader "github.com/bborbe/http/downloader"
	http_downloader_by_url "github.com/bborbe/http/downloader/by_url"
	io_util "github.com/bborbe/io/util"
)

var logger = log.DefaultLogger

const (
	PARAMETER_LOGLEVEL           = "loglevel"
	PARAMETER_PARALLEL_DOWNLOADS = "max"
	PARAMETER_TARGET             = "target"
	DEFAULT_PARALLEL_DOWNLOADS   = 2
	DEFAULT_TARGET               = "~/Downloads"
)

func main() {
	defer logger.Close()
	logLevelPtr := flag.String(PARAMETER_LOGLEVEL, log.INFO_STRING, log.FLAG_USAGE)
	maxConcurrencyDownloadsPtr := flag.Int(PARAMETER_PARALLEL_DOWNLOADS, DEFAULT_PARALLEL_DOWNLOADS, "max parallel downloads")
	targetDirectoryPtr := flag.String(PARAMETER_TARGET, DEFAULT_TARGET, "directory")
	flag.Parse()
	logger.SetLevelThreshold(log.LogStringToLevel(*logLevelPtr))
	logger.Debugf("set log level to %s", *logLevelPtr)

	runtime.GOMAXPROCS(runtime.NumCPU())

	writer := os.Stdout
	input := os.Stdin
	wg := new(sync.WaitGroup)
	httpClientBuilder := http_client_builder.New()
	httpClient := httpClientBuilder.Build()
	downloader := http_downloader_by_url.New(httpClient.Get)

	err := do(writer, input, *maxConcurrencyDownloadsPtr, wg, downloader, *targetDirectoryPtr)
	wg.Wait()
	if err != nil {
		logger.Fatal(err)
		logger.Close()
		os.Exit(1)
	}
}

func do(writer io.Writer, input io.Reader, maxConcurrencyDownloads int, wg *sync.WaitGroup, downloader http_downloader.Downloader, targetDirectoryName string) error {
	var err error
	if targetDirectoryName, err = io_util.NormalizePath(targetDirectoryName); err != nil {
		return err
	}
	if isDir,err := io_util.IsDirectory(targetDirectoryName); err != nil || isDir == false {
		fmt.Fprintf(writer, "parameter %s is invalid\n", PARAMETER_TARGET)
		return fmt.Errorf("parameter is not a directory")
	}
	targetDirectory, err := os.Open(targetDirectoryName)
	if err != nil {
		return err
	}
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
			downloader.Download(link, targetDirectory)
			<-throttle
			wg.Done()
		}()
	}
	return nil
}

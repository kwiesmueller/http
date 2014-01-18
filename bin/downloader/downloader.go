package main

import (
	"bufio"
	"github.com/bborbe/log"
	"os"
)

var logger = log.DefaultLogger

const SIZE = 10

func main() {
	logger.SetLevelThreshold(log.DEBUG)
	defer logger.Close()

	links := make(chan string, SIZE)
	done := make(chan bool, 1)
	go addLinkToChan(links, done)
	for {
		select {
		case link := <-links:
			logger.Debug(link)
		case <-done:
			return
		}
	}
}

func addLinkToChan(links chan<- string, done chan<- bool) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			done <- true
			return
		}
		links <- string(line)
	}
}

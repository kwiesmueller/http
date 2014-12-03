package mock

import (
	"fmt"
	"net/http"

	"github.com/bborbe/log"
)

type requestDownloader struct {
	resp map[string]*http.Response
	err  map[string]error
}

var logger = log.DefaultLogger

func NewRequestDownloaderMock() *requestDownloader {
	d := new(requestDownloader)
	d.resp = make(map[string]*http.Response)
	d.err = make(map[string]error)
	return d
}

func (d *requestDownloader) Download(request *http.Request) (resp *http.Response, err error) {
	url := request.URL.String()
	if d.resp[url] == nil && d.err[url] == nil {
		panic(fmt.Sprintf("no entry found for url '%s'", url))
	}
	return d.resp[url], d.err[url]
}

func (d *requestDownloader) SetDownload(uri string, resp *http.Response, err error) {
	logger.Debugf("SetDownload for url: '%s'", uri)
	fmt.Printf("SetDownload for url: '%s'\n", uri)
	d.resp[uri] = resp
	d.err[uri] = err
}

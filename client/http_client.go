package client

import (
	"io"
	"net/http"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

type RequestDownloader interface {
	Download(request *http.Request) (resp *http.Response, err error)
}

type GetDownloader interface {
	Get(url string) (resp *http.Response, err error)
}

type GetWithHeaderDownloader interface {
	GetWithHeader(url string, header http.Header) (resp *http.Response, err error)
}

type PostDownloader interface {
	Post(url string, body io.Reader) (resp *http.Response, err error)
}

type PostWithHeaderDownloader interface {
	PostWithHeader(url string, header http.Header, body io.Reader) (resp *http.Response, err error)
}

type downloader struct {
	httpClient *http.Client
}

func New(httpClient *http.Client) *downloader {
	d := new(downloader)
	d.httpClient = httpClient
	return d
}

func (d *downloader) Get(url string) (resp *http.Response, err error) {
	return d.BuildRequestAndDownload("GET", url, make(http.Header), nil)
}

func (d *downloader) GetWithHeader(url string, header http.Header) (resp *http.Response, err error) {
	return d.BuildRequestAndDownload("GET", url, header, nil)
}

func (d *downloader) Post(url string, body io.Reader) (resp *http.Response, err error) {
	return d.BuildRequestAndDownload("POST", url, make(http.Header), body)
}

func (d *downloader) PostWithHeader(url string, header http.Header, body io.Reader) (resp *http.Response, err error) {
	return d.BuildRequestAndDownload("POST", url, header, body)
}

func (d *downloader) Download(req *http.Request) (resp *http.Response, err error) {
	logger.Debugf("download %s %v started", req.Method, req.URL)
	defer logger.Debugf("download %s %v finshed", req.Method, req.URL)
	return d.httpClient.Do(req)
}

func (d *downloader) BuildRequestAndDownload(method string, url string, header http.Header, body io.Reader) (resp *http.Response, err error) {
	logger.Debugf("build request for method: %s url: %s", method, url)
	req, err := BuildRequest(method, url, header, body)
	if err != nil {
		return nil, err
	}
	return d.Download(req)
}

func BuildRequest(method string, url string, header http.Header, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header[key] = value
	}
	return req, nil
}

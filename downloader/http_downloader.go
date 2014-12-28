package downloader

import (
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"time"
)

const (
	TIMEOUT             = 30 * time.Second
	KEEPALIVE           = 30 * time.Second
	TLSHANDSHAKETIMEOUT = 10 * time.Second
)

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

type downloader struct{}

func New() *downloader {
	return new(downloader)
}

func (d *downloader) Get(url string) (resp *http.Response, err error) {
	return BuildRequestAndDownload("GET", url, make(http.Header), nil)
}

func (d *downloader) GetWithHeader(url string, header http.Header) (resp *http.Response, err error) {
	return BuildRequestAndDownload("GET", url, header, nil)
}

func (d *downloader) Post(url string, body io.Reader) (resp *http.Response, err error) {
	return BuildRequestAndDownload("POST", url, make(http.Header), body)
}

func (d *downloader) PostWithHeader(url string, header http.Header, body io.Reader) (resp *http.Response, err error) {
	return BuildRequestAndDownload("POST", url, header, body)
}

func (d *downloader) Download(request *http.Request) (resp *http.Response, err error) {
	return Download(request)
}

func BuildRequestAndDownload(method string, url string, header http.Header, body io.Reader) (resp *http.Response, err error) {
	req, err := BuildRequest(method, url, header, body)
	if err != nil {
		return nil, err
	}
	return Download(req)
}

func getClient() *http.Client {
	dialFunc := (&net.Dialer{
		Timeout:   TIMEOUT,
		KeepAlive: KEEPALIVE,
	}).Dial
	tr := &http.Transport{
		Proxy:               http.ProxyFromEnvironment,
		Dial:                dialFunc,
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout: TLSHANDSHAKETIMEOUT,
	}
	client := &http.Client{Transport: tr}
	return client
}

func Download(request *http.Request) (resp *http.Response, err error) {
	client := getClient()
	return client.Do(request)
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
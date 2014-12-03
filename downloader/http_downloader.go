package downloader

import (
	"crypto/tls"
	"io"
	"net/http"
)

type DownloaderGet interface {
	Get(url string) (resp *http.Response, err error)
}

type DownloaderGetWithHeader interface {
	GetWithHeader(url string, header http.Header) (resp *http.Response, err error)
}

type DownloaderPost interface {
	Post(url string, body io.Reader) (resp *http.Response, err error)
}

type DownloaderPostWithHeader interface {
	PostWithHeader(url string, header http.Header, body io.Reader) (resp *http.Response, err error)
}

type downloader struct{}

func New() *downloader {
	return new(downloader)
}

func (d *downloader) Get(url string) (resp *http.Response, err error) {
	return Download("GET", url, make(http.Header), nil)
}

func (d *downloader) GetWithHeader(url string, header http.Header) (resp *http.Response, err error) {
	return Download("GET", url, header, nil)
}

func (d *downloader) Post(url string, body io.Reader) (resp *http.Response, err error) {
	return Download("POST", url, make(http.Header), body)
}

func (d *downloader) PostWithHeader(url string, header http.Header, body io.Reader) (resp *http.Response, err error) {
	return Download("POST", url, header, body)
}

func getClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	return client
}

func Download(method string, url string, header http.Header, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for key, value := range header {
		req.Header[key] = value
	}
	client := getClient()
	return client.Do(req)
}

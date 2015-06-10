package mock

import (
	"io"
	"net/http"
)

type GetDownloaderMock struct{}

func NewGetDownloader() *GetDownloaderMock {
	return new(GetDownloaderMock)
}

type PostDownloaderMock struct{}

func NewPostDownloader() *PostDownloaderMock {
	return new(PostDownloaderMock)
}

type GetWithHeaderDownloaderMock struct{}

func NewGetWithHeaderDownloader() *GetWithHeaderDownloaderMock {
	return new(GetWithHeaderDownloaderMock)
}

type PostWithHeaderDownloaderMock struct{}

func NewPostWithHeaderDownloader() *PostWithHeaderDownloaderMock {
	return new(PostWithHeaderDownloaderMock)
}

type RequestDownloaderMock struct {
	Request  *http.Request
	Response *http.Response
	Error    error
}

func NewRequestDownloader() *RequestDownloaderMock {
	return new(RequestDownloaderMock)
}

func (o *RequestDownloaderMock) Download(request *http.Request) (resp *http.Response, err error) {
	o.Request = request
	return o.Response, o.Error
}

func (o *GetDownloaderMock) Get(url string) (resp *http.Response, err error) {
	return nil, nil
}

func (o *GetWithHeaderDownloaderMock) GetWithHeader(url string, header http.Header) (resp *http.Response, err error) {
	return nil, nil
}

func (o *PostDownloaderMock) Post(url string, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}

func (o *PostWithHeaderDownloaderMock) PostWithHeader(url string, header http.Header, body io.Reader) (resp *http.Response, err error) {
	return nil, nil
}

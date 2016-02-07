package client_builder

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"net/url"

	"errors"

	"github.com/bborbe/log"
)

var logger = log.DefaultLogger

const (
	TIMEOUT = 30 * time.Second
	KEEPALIVE = 30 * time.Second
	TLSHANDSHAKETIMEOUT = 10 * time.Second
)

type HttpClientBuilder interface {
	Build() *http.Client
	BuildRoundTripper() http.RoundTripper
	WithProxy() HttpClientBuilder
	WithoutProxy() HttpClientBuilder
	WithRedirects() HttpClientBuilder
	WithoutRedirects() HttpClientBuilder
}

type httpClientBuilder struct {
	proxy         Proxy
	checkRedirect CheckRedirect
}

type Proxy func(req *http.Request) (*url.URL, error)

type CheckRedirect func(req *http.Request, via []*http.Request) error

func New() *httpClientBuilder {
	b := new(httpClientBuilder)
	b.WithoutProxy()
	b.WithRedirects()
	return b
}

func (b *httpClientBuilder) BuildRoundTripper() http.RoundTripper {
	logger.Debugf("build http transport")
	dialFunc := (&net.Dialer{
		Timeout: TIMEOUT,
		//		KeepAlive: KEEPALIVE,
	}).Dial
	return &http.Transport{
		Proxy:           b.proxy,
		Dial:            dialFunc,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		//		TLSHandshakeTimeout: TLSHANDSHAKETIMEOUT,
	}
}

func (b *httpClientBuilder) Build() *http.Client {
	logger.Debugf("build http client")
	return &http.Client{
		Transport:     b.BuildRoundTripper(),
		CheckRedirect: b.checkRedirect,
	}
}

func (b *httpClientBuilder) WithProxy() HttpClientBuilder {
	b.proxy = http.ProxyFromEnvironment
	return b
}

func (b *httpClientBuilder) WithoutProxy() HttpClientBuilder {
	b.proxy = nil
	return b
}

func (b *httpClientBuilder) WithRedirects() HttpClientBuilder {
	b.checkRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= 10 {
			return errors.New("stopped after 10 redirects")
		}
		return nil
	}
	return b
}

func (b *httpClientBuilder) WithoutRedirects() HttpClientBuilder {
	b.checkRedirect = func(req *http.Request, via []*http.Request) error {
		if len(via) >= 1 {
			return errors.New("redirects")
		}
		return nil
	}
	return b
}

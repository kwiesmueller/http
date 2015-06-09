package requestbuilder

import (
	"io"
	"net/http"
)

type HttpRequestBuilder interface {
	AddParameter(key string, value ...string)
	AddHeader(key string, values ...string)
	SetMethod(key string)
	SetBody(reader io.Reader)
	GetRequest() (*http.Request, error)
}

type httpRequestBuilder struct {
	url       string
	parameter map[string][]string
	header    http.Header
	method    string
	body      io.Reader
}

func NewHttpRequestBuilder(url string) *httpRequestBuilder {
	r := new(httpRequestBuilder)
	r.method = "GET"
	r.url = url
	r.parameter = make(map[string][]string)
	r.header = make(http.Header)
	return r
}

func (r *httpRequestBuilder) SetBody(body io.Reader) {
	r.body = body
}

func (r *httpRequestBuilder) SetMethod(method string) {
	r.method = method
}

func (r *httpRequestBuilder) AddHeader(key string, values ...string) {
	r.header[key] = values
}

func (r *httpRequestBuilder) AddParameter(key string, values ...string) {
	r.parameter[key] = values
}

func (r *httpRequestBuilder) GetRequest() (*http.Request, error) {
	req, err := http.NewRequest(r.method, r.getUrlWithParameter(), r.body)
	if err != nil {
		return nil, err
	}
	req.Header = r.header
	return req, nil
}

func (r *httpRequestBuilder) getUrlWithParameter() string {
	result := r.url
	first := true
	for key, values := range r.parameter {
		for _, value := range values {
			if first {
				first = false
				result += "?"
			} else {
				result += "&"
			}
			result += key
			result += "="
			result += value
		}
	}
	return result
}

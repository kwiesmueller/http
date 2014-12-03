package requestbuilder

type HttpRequestBuilderProvider interface {
	NewHttpRequestBuilder(url string) HttpRequestBuilder
}

type httpRequestBuilderProvider struct {
}

func NewHttpRequestBuilderProvider() *httpRequestBuilderProvider {
	p := new(httpRequestBuilderProvider)
	return p
}

func (p *httpRequestBuilderProvider) NewHttpRequestBuilder(url string) HttpRequestBuilder {
	return NewHttpRequestBuilder(url)
}

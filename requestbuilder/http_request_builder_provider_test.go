package requestbuilder

import (
	"testing"
	. "github.com/bborbe/assert"
)

func TestImplementsNewHttpRequestBuilderProvider(t *testing.T) {
	p := NewHttpRequestBuilderProvider()
	var i *HttpRequestBuilderProvider
	err := AssertThat(p, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewHttpRequestBuilder(t *testing.T) {
	var err error
	p := NewHttpRequestBuilderProvider()
	rb := p.NewHttpRequestBuilder("http://example.com")
	err = AssertThat(rb, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	var i *HttpRequestBuilder
	err = AssertThat(rb, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

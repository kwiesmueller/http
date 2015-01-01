package requestbuilder

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestImplementsHttpRequestBuilder(t *testing.T) {
	r := NewHttpRequestBuilder("http://www.example.com")
	var i *HttpRequestBuilder
	err := AssertThat(r, Implements(i))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetRequestWithHeader(t *testing.T) {
	r := NewHttpRequestBuilder("http://www.benjamin-borbe.de")
	r.AddHeader("a", "b")
	request, err := r.GetRequest()
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(request, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(request.Header), Is(1))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(request.Header["a"]), Is(1))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(request.Header["a"][0], Is("b"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetRequest(t *testing.T) {
	r := NewHttpRequestBuilder("http://www.benjamin-borbe.de")
	request, err := r.GetRequest()
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(request, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
}

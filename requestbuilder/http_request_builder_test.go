package requestbuilder

import (
	"io/ioutil"
	"net/http"
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

func TestGetResponse(t *testing.T) {
	r := NewHttpRequestBuilder("http://www.benjamin-borbe.de")
	r.AddParameter("a", "b")
	response, err := r.GetResponse()
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(response, NotNilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(response.StatusCode, Is(http.StatusOK))
	if err != nil {
		t.Fatal(err)
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(len(content), Gt(0))
	if err != nil {
		t.Fatal(err)
	}
}

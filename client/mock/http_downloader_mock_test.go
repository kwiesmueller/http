package mock

import (
	"testing"

	. "github.com/bborbe/assert"
	http_client "github.com/bborbe/http/client"
 	http_client_builder "github.com/bborbe/http/client/builder"
 )

func TestImplementsDownloaderGet(t *testing.T) {
	d := NewGetDownloader()
	var i *http_client.GetDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderPost(t *testing.T) {
	d := NewPostDownloader()
	var i *http_client.PostDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderGetWithHeader(t *testing.T) {
	d := NewGetWithHeaderDownloader()
	var i *http_client.GetWithHeaderDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderPostWithHeader(t *testing.T) {
	d := NewPostWithHeaderDownloader()
	var i *http_client.PostWithHeaderDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsRequestDownloader(t *testing.T) {
	d := NewRequestDownloader()
	var i *http_client.RequestDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

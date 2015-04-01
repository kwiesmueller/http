package mock

import (
	"testing"

	. "github.com/bborbe/assert"
	"github.com/bborbe/http/client"
)

func TestImplementsDownloaderGet(t *testing.T) {
	d := NewGetDownloader()
	var i *client.GetDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderPost(t *testing.T) {
	d := NewPostDownloader()
	var i *client.PostDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderGetWithHeader(t *testing.T) {
	d := NewGetWithHeaderDownloader()
	var i *client.GetWithHeaderDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderPostWithHeader(t *testing.T) {
	d := NewPostWithHeaderDownloader()
	var i *client.PostWithHeaderDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsRequestDownloader(t *testing.T) {
	d := NewRequestDownloader()
	var i *client.RequestDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

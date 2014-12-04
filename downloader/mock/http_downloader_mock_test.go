package mock

import (
	"testing"
	. "github.com/bborbe/assert"
	"github.com/bborbe/http/downloader"
)

func TestImplementsDownloaderGet(t *testing.T) {
	d := NewGetDownloader()
	var i *downloader.GetDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderPost(t *testing.T) {
	d := NewPostDownloader()
	var i *downloader.PostDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderGetWithHeader(t *testing.T) {
	d := NewGetWithHeaderDownloader()
	var i *downloader.GetWithHeaderDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderPostWithHeader(t *testing.T) {
	d := NewPostWithHeaderDownloader()
	var i *downloader.PostWithHeaderDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsRequestDownloader(t *testing.T) {
	d := NewRequestDownloader()
	var i *downloader.RequestDownloader
	err := AssertThat(d, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

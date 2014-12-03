package downloader

import (
	"testing"
	. "github.com/bborbe/assert"
)

func TestImplementsDownloaderGet(t *testing.T) {
	downloader := New()
	var i *GetDownloader
	err := AssertThat(downloader, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderPost(t *testing.T) {
	downloader := New()
	var i *PostDownloader
	err := AssertThat(downloader, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderGetWithHeader(t *testing.T) {
	downloader := New()
	var i *GetWithHeaderDownloader
	err := AssertThat(downloader, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderPostWithHeader(t *testing.T) {
	downloader := New()
	var i *PostWithHeaderDownloader
	err := AssertThat(downloader, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsRequestDownloader(t *testing.T) {
	downloader := New()
	var i *RequestDownloader
	err := AssertThat(downloader, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

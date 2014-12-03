package downloader

import (
	"testing"
	. "github.com/bborbe/assert"
)

func TestImplementsDownloaderGet(t *testing.T) {
	downloader := New()
	var i *DownloaderGet
	err := AssertThat(downloader, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderPost(t *testing.T) {
	downloader := New()
	var i *DownloaderPost
	err := AssertThat(downloader, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderGetWithHeader(t *testing.T) {
	downloader := New()
	var i *DownloaderGetWithHeader
	err := AssertThat(downloader, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderPostWithHeader(t *testing.T) {
	downloader := New()
	var i *DownloaderPostWithHeader
	err := AssertThat(downloader, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestImplementsDownloaderRequest(t *testing.T) {
	downloader := New()
	var i *DownloaderRequest
	err := AssertThat(downloader, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

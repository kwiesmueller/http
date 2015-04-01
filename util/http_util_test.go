package util

import (
	"net/http"
	"testing"

	"net/url"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
)

func TestResponseToByteArray(t *testing.T) {
	var err error
	var content []byte

	response := new(http.Response)
	response.Body = io_mock.NewReadCloserString("test")

	if content, err = ResponseToByteArray(response); err != nil {
		t.Fatal(err)
	}

	if err = AssertThat(string(content), Is("test")); err != nil {
		t.Fatal(err)
	}
}

func TestResponseToString(t *testing.T) {
	var err error
	var content string

	response := new(http.Response)
	response.Body = io_mock.NewReadCloserString("test")

	if content, err = ResponseToString(response); err != nil {
		t.Fatal(err)
	}

	if err = AssertThat(content, Is("test")); err != nil {
		t.Fatal(err)
	}
}

func TestFindFileExtension(t *testing.T) {
	var err error

	response := &http.Response{}

	if err = AssertThat(FindFileExtension(response), Is("")); err != nil {
		t.Fatal(err)
	}
}

func TestFindFileExtensionUrlWithDot(t *testing.T) {
	var err error
	var u *url.URL
	if u, err = url.ParseRequestURI("http://www.example/robots.txt"); err != nil {
		t.Fatal(err)
	}
	response := &http.Response{Request: &http.Request{URL: u}}

	if err = AssertThat(FindFileExtension(response), Is("txt")); err != nil {
		t.Fatal(err)
	}
}

func TestFindFileExtensionUrlWithDotAtLast(t *testing.T) {
	var err error
	var u *url.URL
	if u, err = url.ParseRequestURI("http://www.example/robots."); err != nil {
		t.Fatal(err)
	}
	response := &http.Response{Request: &http.Request{URL: u}}

	if err = AssertThat(FindFileExtension(response), Is("")); err != nil {
		t.Fatal(err)
	}
}

func TestFindFileExtensionHeader(t *testing.T) {
	var err error
	response := &http.Response{Header: http.Header{}}
	if err = AssertThat(FindFileExtension(response), Is("")); err != nil {
		t.Fatal(err)
	}
}

func TestFindFileExtensionHeaderContentTypeKownType(t *testing.T) {
	var err error
	response := &http.Response{Header: http.Header{"Content-Type": []string{"image/jpeg"}}}
	if err = AssertThat(FindFileExtension(response), Is("jpg")); err != nil {
		t.Fatal(err)
	}
}

func TestFindFileExtensionHeaderContentTypeUnkownType(t *testing.T) {
	var err error
	response := &http.Response{Header: http.Header{"Content-Type": []string{"text/foo"}}}
	if err = AssertThat(FindFileExtension(response), Is("")); err != nil {
		t.Fatal(err)
	}
}

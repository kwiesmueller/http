package http

import (
	"net/http"
	"testing"
	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
)

func TestResponseToByteArray(t *testing.T) {
	response := new(http.Response)
	response.Body = io_mock.NewReadCloserString("test")
	content, err := ResponseToByteArray(response)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(content), Is("test"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestResponseToString(t *testing.T) {
	response := new(http.Response)
	response.Body = io_mock.NewReadCloserString("test")
	content, err := ResponseToString(response)
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(content, Is("test"))
	if err != nil {
		t.Fatal(err)
	}
}

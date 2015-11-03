package main

import (
	"testing"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
)

func TestDo(t *testing.T) {
	var err error
	writer := io_mock.NewWriter()
	input := io_mock.NewReadCloserString(" http://www.example.com ")
	err = do(writer, input)
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("http://www.example.com\n"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoBugfix(t *testing.T) {
	var err error
	writer := io_mock.NewWriter()
	input := io_mock.NewReadCloserString("\n2015-04-01T15:40:09 http://www.example.com\n2015-04-01T15:40:09 http://www.example.com\n")
	err = do(writer, input)
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("http://www.example.com\nhttp://www.example.com\n"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestDoWithParameters(t *testing.T) {
	var err error
	writer := io_mock.NewWriter()
	input := io_mock.NewReadCloserString(" http://www.example.com?a=b ")
	err = do(writer, input)
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("http://www.example.com?a=b\n"))
	if err != nil {
		t.Fatal(err)
	}
}

package main

import (
	"testing"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
)

func TestDo(t *testing.T) {
	var err error
	writer := io_mock.NewWriter()
	input := io_mock.NewReadCloserString("  http://www.exmaple.com  ")
	err = do(writer, input)
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(string(writer.Content()), Is("http://www.exmaple.com\n"))
	if err != nil {
		t.Fatal(err)
	}
}

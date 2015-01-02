package main

import (
	"testing"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
)

func TestDo(t *testing.T) {
	writer := io_mock.NewWriter()
	input := io_mock.NewReadCloserString("")
	err := do(writer, input)
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
}

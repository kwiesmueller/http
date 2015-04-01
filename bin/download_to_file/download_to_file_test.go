package main

import (
	"testing"

	"sync"

	. "github.com/bborbe/assert"
	io_mock "github.com/bborbe/io/mock"
)

func TestDo(t *testing.T) {
	writer := io_mock.NewWriter()
	input := io_mock.NewReadCloserString("")
	wg := new(sync.WaitGroup)
	err := do(writer, input, 2, wg, nil, "/tmp")
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
}

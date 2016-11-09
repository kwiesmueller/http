package rest

import (
	"testing"

	"os"

	. "github.com/bborbe/assert"
	"github.com/golang/glog"
)

func TestMain(m *testing.M) {
	exit := m.Run()
	glog.Flush()
	os.Exit(exit)
}

func TestImplementsRest(t *testing.T) {
	c := New(nil)
	var i *Rest
	if err := AssertThat(c, Implements(i)); err != nil {
		t.Fatal(err)
	}
}

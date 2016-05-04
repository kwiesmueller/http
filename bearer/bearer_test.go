package bearer

import (
	"testing"

	. "github.com/bborbe/assert"
)

func TestCreateParseBearer(t *testing.T) {
	header := CreateBearerHeader("foo", "bar")
	name, value, err := ParseBearerHeader(header)
	if err := AssertThat(err, NilValue()); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(name, Is("foo")); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(value, Is("bar")); err != nil {
		t.Fatal(err)
	}
}

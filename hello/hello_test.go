package hello_test

import (
	"bytes"
	"testing"

	"github.com/wolves/hello"
)

func TestPrintTo_PrintsHelloMessageToGivenWriter(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	hello.PrintTo(buf)
	want := "Hello, world\n"
	got := buf.String()
	if got != want {
		t.Errorf("want %q, got %q", want, got)
	}
}

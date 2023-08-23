package hello_test

import (
	"bytes"
	"testing"

	"github.com/wolves/hello"
)

func TestPrint_PrintsHelloMessageToOutput(t *testing.T) {
	t.Parallel()
	buf := new(bytes.Buffer)
	p := hello.NewPrinter()
	p.Output = buf
	p.Print()
	want := "Hello, world\n"
	got := buf.String()
	if want != got {
		t.Errorf("want %q, but got %q", want, got)
	}
}

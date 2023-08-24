package count_test

import (
	"bytes"
	"testing"

	"github.com/wolves/count"
)

func TestLines_CountsLinesInInput(t *testing.T) {
	lines := "Line 1\nLine 2\nLine 3"

	t.Parallel()
	c := count.NewCounter()
	c.Input = bytes.NewBufferString(lines)
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

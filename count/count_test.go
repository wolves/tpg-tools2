package count_test

import (
	"bytes"
	"testing"

	"github.com/wolves/count"
)

func TestLines_CountsLinesInInput(t *testing.T) {
	t.Parallel()
	lines := "Line 1\nLine 2\nLine 3"
	inBuf := bytes.NewBufferString(lines)
	c, err := count.NewCounter(
		count.WithInput(inBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

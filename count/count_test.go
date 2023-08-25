package count_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"
	"github.com/wolves/count"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"count": count.Main,
	}))
}

func Test(t *testing.T) {
	t.Parallel()
	testscript.Run(t, testscript.Params{
		Dir: "testdata/script",
	})
}

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

func TestWords_CountsWordsInInput(t *testing.T) {
	t.Parallel()
	inBuf := bytes.NewBufferString("1\n2 words\n3 this time")
	c, err := count.NewCounter(
		count.WithInput(inBuf),
	)
	if err != nil {
		t.Fatal(err)
	}

	want := 6
	got := c.Words()

	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBytes_CountsBytesInInput(t *testing.T) {
	t.Parallel()
	inBuf := bytes.NewBufferString("1\n2 words\n3 this time")
	c, err := count.NewCounter(
		count.WithInput(inBuf),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 21
	got := c.Bytes()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs_SetsInputToGivenPath(t *testing.T) {
	t.Parallel()
	args := []string{"testdata/three_lines.txt"}
	c, err := count.NewCounter(
		count.WithInputFromArgs(args),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestWithInputFromArgs_UsesDefaultInputWithNoArgs(t *testing.T) {
	t.Parallel()
	c, err := count.NewCounter(
		count.WithInput(bytes.NewBufferString("Line 1\nLine 2\nLine 3")),
		count.WithInputFromArgs([]string{}),
	)
	if err != nil {
		t.Fatal(err)
	}
	want := 3
	got := c.Lines()
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

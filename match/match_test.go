package match_test

import (
	"bytes"
	"testing"

	"github.com/wolves/match"
)

func TestMatchingLines_PrintsLinesWithArgMatchToOutput(t *testing.T) {
	t.Parallel()
	inBuf := bytes.NewBufferString("match line 1\nregular line 2\nanothermatch 3")
	outBuf := new(bytes.Buffer)
	searchArgs := []string{"match"}

	m, err := match.NewMatcher(
		match.WithInput(inBuf),
		match.WithOutput(outBuf),
		match.WithArgsSearchText(searchArgs),
	)
	if err != nil {
		t.Fatal(err)
	}

	m.MatchingLines()

	want := "match line 1\nanothermatch 3\n"
	got := outBuf.String()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}

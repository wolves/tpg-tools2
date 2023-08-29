package findgo_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wolves/findgo"
)

func TestFilesCorrectylListsFileInTree(t *testing.T) {
	t.Parallel()
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	got := findgo.Files("testdata/tree")
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

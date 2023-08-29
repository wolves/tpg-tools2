package findgo_test

import (
	"archive/zip"
	"os"
	"testing"
	"testing/fstest"

	"github.com/google/go-cmp/cmp"
	"github.com/wolves/findgo"
)

func BenchmarkFilesOnDisk(b *testing.B) {
	fsys := os.DirFS("testing/tree")
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findgo.Files(fsys)
	}
}

func BenchmarkFilesInMemory(b *testing.B) {
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = findgo.Files(fsys)
	}
}

func TestFilesCorrectylListsFileInTree(t *testing.T) {
	t.Parallel()
	fsys := os.DirFS("testdata/tree")
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	got := findgo.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFilesCorrectlyListsFilesInMapFS(t *testing.T) {
	t.Parallel()
	fsys := fstest.MapFS{
		"file.go":                {},
		"subfolder/subfolder.go": {},
		"subfolder2/another.go":  {},
		"subfolder2/file.go":     {},
	}
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	got := findgo.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

func TestFilesCorrectlyListsFilesInZIPArchive(t *testing.T) {
	t.Parallel()
	fsys, err := zip.OpenReader("testdata/files.zip")
	if err != nil {
		t.Fatal(err)
	}
	want := []string{
		"file.go",
		"subfolder/subfolder.go",
		"subfolder2/another.go",
		"subfolder2/file.go",
	}
	got := findgo.Files(fsys)
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}
}

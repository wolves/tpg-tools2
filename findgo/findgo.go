package findgo

import (
	"io/fs"
	"os"
	"path/filepath"
)

func Files(p string) []string {
	files := []string{}
	fsys := os.DirFS(p)
	fs.WalkDir(fsys, ".", func(p string, d fs.DirEntry, err error) error {
		if filepath.Ext(p) == ".go" {
			files = append(files, p)
		}
		return nil
	})
	return files
}

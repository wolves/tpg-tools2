package older

import (
	"io/fs"
	"time"
)

func Find(fsys fs.FS, age time.Duration) (paths []string) {
	fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		cutoff := time.Now().Add(-age)
		info, err := d.Info()
		if err != nil || info.IsDir() {
			return nil
		}

		if info.ModTime().Before(cutoff) {
			paths = append(paths, path)
		}

		return nil
	})
	return paths
}

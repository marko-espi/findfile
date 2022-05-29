package findfile

import (
	"errors"
	"io/fs"
	"path/filepath"
)

func File(fsys fs.FS, ext string) (count int, err error) {
	if ext == "" {
		return 0, errors.New("extention cannot be blank")
	}
	err = fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if filepath.Ext(path) == ext {
			count++
		}
		return err
	})
	if err != nil {
		return 0, err
	}
	return count, nil
}

package generator

import (
	"errors"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

var ErrEmptyTemplate = errors.New("template name cannot be empty")

func CopyFile(src fs.FS, srcPath, dstPath string) error {
	in, err := src.Open(srcPath)
	if err != nil {
		return err
	}
	defer in.Close()

	if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
		return err
	}

	out, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	return err
}

func TemplateExists(fsys fs.FS, name string) (bool, error) {
	if name == "" {
		return false, ErrEmptyTemplate
	}

	_, err := fs.Stat(fsys, path.Join("templates", name))
	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	return false, err
}

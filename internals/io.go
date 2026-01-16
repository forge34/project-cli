package internals

import (
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

func CopyFileWithTemplate(
	src fs.FS,
	relPath string,
	dstPath string,
	vars map[string]string,
) error {
	tmpl, err := template.ParseFS(src, relPath)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(dstPath), 0o755); err != nil {
		return err
	}

	out, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer out.Close()

	return tmpl.Execute(out, vars)
}

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

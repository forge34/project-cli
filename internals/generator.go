// Package internals provides functionality for creating projects from templates.
package internals

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// Templates Embed the entire directory.
//
//go:embed templates/*
var Templates embed.FS

type (
	Generator struct{}
)

func (g *Generator) Create(fsys fs.FS, dst string) error {
	prompts, err := ParseTemplate(fsys)
	if err != nil {
		return err
	}

	var answers map[string]string
	if len(prompts.Prompts) >= 1 {
		answers, err = promptUser(prompts)
	}

	if err != nil {
		return err
	}

	err = fs.WalkDir(fsys, ".", func(rel string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}

		if d.Name() == "template.json" {
			return nil
		}

		if _, err := os.Stat(target); err == nil {
			fmt.Printf("%s already exists,Skipping \n", rel)
			return nil
		}

		if strings.HasSuffix(rel, ".tmpl") {
			target = strings.TrimSuffix(target, ".tmpl")
			return CopyFileWithTemplate(fsys, rel, target, answers)
		}

		return CopyFile(fsys, rel, target)
	})
	if err != nil {
		return err
	}
	return nil
}

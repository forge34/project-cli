// Package generator provides functionality for creating projects from templates.
package generator

import (
	"embed"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
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

var ErrNotExist = errors.New("template doesn't exist")

func (g *Generator) Create(tempName string, dst string) error {
	found, err := TemplateExists(Templates, tempName)
	if err != nil {
		return err
	}

	if !found {
		return ErrNotExist
	}

	base := path.Join("templates", tempName)

	sub, err := fs.Sub(Templates, base)
	if err != nil {
		return err
	}

	prompts, err := ParseTemplate(sub, tempName)
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

	err = fs.WalkDir(sub, ".", func(rel string, d fs.DirEntry, err error) error {
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
			return CopyFileWithTemplate(sub, rel, target, answers)
		}

		return CopyFile(sub, rel, target)
	})
	if err != nil {
		return err
	}
	return nil
}

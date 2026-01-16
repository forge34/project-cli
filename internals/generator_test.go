package internals_test

import (
	"io/fs"
	"path"
	"path/filepath"
	"testing"

	generator "pjc/internals"
)

func TestGeneratesDir(t *testing.T) {
	tmpRoot := t.TempDir()
	dst := filepath.Join(tmpRoot, "server")
	base := path.Join("templates", "express-ts")

	sub, _ := fs.Sub(generator.Templates, base)

	g := generator.Generator{}
	if err := g.Create(sub, dst); err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
}

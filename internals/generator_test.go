package generator_test

import (
	"path/filepath"
	"testing"

	generator "pjc/internals"
)

func TestGeneratesDir(t *testing.T) {
	template := "express-ts"

	tmpRoot := t.TempDir()
	dst := filepath.Join(tmpRoot, "server")

	g := generator.Generator{}
	if err := g.Create(template, dst); err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
}

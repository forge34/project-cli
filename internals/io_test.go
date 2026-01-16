package internals

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestCopyFile(t *testing.T) {
	srcFS := fstest.MapFS{
		"file.txt": {Data: []byte("hello world")},
	}

	tmpDir := t.TempDir()
	dstPath := filepath.Join(tmpDir, "copied.txt")

	err := CopyFile(srcFS, "file.txt", dstPath)
	if err != nil {
		t.Fatalf("copyFile failed: %v", err)
	}

	data, err := os.ReadFile(dstPath)
	if err != nil {
		t.Fatalf("reading copied file failed: %v", err)
	}

	if string(data) != "hello world" {
		t.Errorf("expected 'hello world', got %q", string(data))
	}
}

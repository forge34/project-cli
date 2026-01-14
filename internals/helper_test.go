package generator

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

func TestTemplateExists(t *testing.T) {
	fs := fstest.MapFS{
		"templates/web/index.html":  {Data: []byte("html")},
		"templates/web/style.css":   {Data: []byte("css")},
		"templates/cli/config.yaml": {Data: []byte("yaml")},
		"templates/empty/.keep":     {Data: []byte("")},
	}

	tests := []struct {
		name     string
		tempName string
		want     bool
		wantErr  bool
	}{
		{"empty name", "", false, true},
		{"existing template web", "web", true, false},
		{"existing template cli", "cli", true, false},
		{"existing empty template", "empty", true, false},
		{"nonexistent template", "nope", false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := TemplateExists(fs, tt.tempName)

			if (err != nil) != tt.wantErr {
				t.Fatalf("TemplateExists() error = %v, wantErr %v", err, tt.wantErr)
			}

			if got != tt.want {
				t.Errorf("TemplateExists() = %v, want %v", got, tt.want)
			}
		})
	}
}

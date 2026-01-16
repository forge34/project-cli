package internals

import (
	"testing"
	"testing/fstest"
)

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

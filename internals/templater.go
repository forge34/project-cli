package generator

import (
	"encoding/json"
	"errors"
	"io/fs"
)

type TemplateVar struct {
	Name    string `json:"name"`
	Prompt  string `json:"prompt"`
	Default string `json:"default,omitempty"`
}

type TemplateConfig struct {
	Prompts []TemplateVar `json:"prompts"`
}

var (
	ErrEmptyTemplateName = errors.New("template name cannot be empty")
	ErrTemplateNotFound  = errors.New("template.json not found")
)

func ParseTemplate(fsys fs.FS, name string) (TemplateConfig, error) {
	if name == "" {
		return TemplateConfig{}, ErrEmptyTemplateName
	}
	buf, err := fs.ReadFile(fsys, "template.json")
	if err != nil {
		return TemplateConfig{}, ErrTemplateNotFound
	}

	var config TemplateConfig
	if err := json.Unmarshal(buf, &config); err != nil {
		return TemplateConfig{}, err
	}

	return config, nil
}

package generator

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestParseTemplate(t *testing.T) {
	fsys := fstest.MapFS{
		"my-template/template.json": {
			Data: []byte(`
{
  "prompts": [
    {
      "name": "projectName",
      "prompt": "Project name",
      "default": "my-app"
    },
    {
      "name": "author",
      "prompt": "Author name"
    }
  ]
}
`),
		},
	}

	cfg, err := ParseTemplate(fsys, "my-template")
	if err != nil {
		t.Fatalf("ParseTemplate returned error: %v", err)
	}

	expected := TemplateConfig{
		Prompts: []TemplateVar{
			{
				Name:    "projectName",
				Prompt:  "Project name",
				Default: "my-app",
			},
			{
				Name:   "author",
				Prompt: "Author name",
			},
		},
	}

	if !reflect.DeepEqual(cfg, expected) {
		t.Errorf("ParseTemplate() result mismatch\nexpected: %#v\ngot:      %#v", expected, cfg)
	}
}

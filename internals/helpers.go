package internals

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"strings"
)

var ErrEmptyTemplate = errors.New("template name cannot be empty")

func promptUser(prompts TemplateConfig) (map[string]string, error) {
	reader := bufio.NewReader(os.Stdin)
	ask := func(label string) (string, error) {
		fmt.Print(Success.Render("? ") + Header.Render(label) + " ")
		s, err := reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		return strings.TrimSpace(s), nil
	}
	answers := make(map[string]string)
	for _, v := range prompts.Prompts {
		ans, err := ask(v.Prompt)
		if err != nil {
			return make(map[string]string), nil
		}

		answers[v.Name] = ans
	}

	return answers, nil
}

func TemplateExists(fsys fs.FS, name string) (bool, error) {
	if name == "" {
		return false, ErrEmptyTemplate
	}

	_, err := fs.Stat(fsys, path.Join("templates", name))
	if err == nil {
		return true, nil
	}

	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}

	return false, err
}

func ListTemplates(fsys fs.FS) ([]string, error) {
	var templateList []string
	maxDepth := 2
	err := fs.WalkDir(fsys, ".", func(rel string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		depth := strings.Count(rel, string(os.PathSeparator))

		if d.Name() == "." {
			return nil
		}

		if depth >= maxDepth {
			return fs.SkipDir
		}

		if d.IsDir() && depth == 1 {
			templateList = append(templateList, d.Name())
		}

		return nil
	})
	if err != nil {
		return []string{}, err
	}

	return templateList, nil
}

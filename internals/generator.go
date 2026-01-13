// Package generator provides functionality for creating projects from templates.
package generator

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

type Generator struct{}

func (g *Generator) Generate(tempName string, dir string) {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	fullPath := filepath.Join(pwd, dir)

	fmt.Println(fullPath)
	err = filepath.WalkDir(filepath.Join("./templates", tempName), func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		fmt.Println(d)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

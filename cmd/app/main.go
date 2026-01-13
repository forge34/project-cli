package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	generator "pjc/internals"
)

func templateExists(tempName string) bool {
	found := false

	if tempName == "" {
		fmt.Println("Template can't be empty")
		os.Exit(1)
	}

	err := filepath.WalkDir("./templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.Name() == tempName {
			found = true
			return filepath.SkipAll
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return found
}

func main() {
	template := flag.String("template", "", "The template used for the project")
	flag.Parse()

	if !templateExists(*template) {
		fmt.Println("Template not found")
		os.Exit(1)
	}

	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("No command provided")
		os.Exit(1)
	}

	cmd := args[0]

	g := &generator.Generator{}
	switch cmd {
	case "create":
		if len(args) <= 1 || args[1] == "" {
			fmt.Println("Directory name can't be empty")
			os.Exit(1)
		}

		dirName := args[1]
		fmt.Println("Creating project in directory:", dirName)
		g.Generate(*template, dirName)
		os.Exit(0)

	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}

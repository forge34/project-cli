package main

import (
	"flag"
	"fmt"
	"os"

	generator "pjc/internals"
)

func main() {
	template := flag.String("template", "", "The template used for the project")
	flag.Parse()

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
		if err := g.Generate(*template, dirName); err != nil {
			panic(err)
		}
		os.Exit(0)
	case "list":
		if err := generator.ListTempaltes(generator.Templates); err != nil {
			panic(err)
		}
		os.Exit(0)

	default:
		fmt.Println("Invalid command")
		os.Exit(1)
	}
}

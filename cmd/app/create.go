package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	generator "pjc/internals"

	"github.com/spf13/cobra"
)

var ErrNotExist = errors.New("template doesn't exist")

var createCmd = &cobra.Command{
	Use:   "create <template> <destination>",
	Short: "Generate a new project from a template",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		templateName := args[0]
		dstArg := args[1]
		found, err := generator.TemplateExists(generator.Templates, templateName)
		if err != nil {
			return err
		}

		if !found {
			return ErrNotExist
		}
		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		dst := filepath.Join(pwd, dstArg)

		g := generator.Generator{}
		base := path.Join("templates", templateName)

		sub, err := fs.Sub(generator.Templates, base)
		if err != nil {
			return err
		}
		err = g.Create(sub, dst)
		if err != nil {
			return err
		}

		fmt.Printf("Project generated at %s\n", dst)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(createCmd)
}

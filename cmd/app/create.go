package main

import (
	"fmt"
	"os"
	"path/filepath"

	generator "pjc/internals"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create <template> <destination>",
	Short: "Generate a new project from a template",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		templateName := args[0]
		dstArg := args[1]

		pwd, err := os.Getwd()
		if err != nil {
			return err
		}

		dst := filepath.Join(pwd, dstArg)

		g := generator.Generator{}
		err = g.Generate(templateName, dst)
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

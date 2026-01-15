package main

import (
	generator "pjc/internals"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := generator.ListTempaltes(generator.Templates); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

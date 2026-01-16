package main

import (
	"pjc/internals"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all templates",
	RunE: func(cmd *cobra.Command, args []string) error {
		l, err := internals.ListTemplates(internals.Templates)
		if err != nil {
			return err
		}

		internals.PrintList(l)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}

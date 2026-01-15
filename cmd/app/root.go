package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "pjc",
	Short: "Project generator CLI",
	Long:  "Generate projects from embedded templates with variable substitution",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Use 'pjc init <template> <destination>' or 'pjc help'")
	},
}

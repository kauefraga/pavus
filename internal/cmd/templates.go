package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kauefraga/pavus/internal/markdown"
	"github.com/spf13/cobra"
)

func getTemplatesCmd() *cobra.Command {
	templatesCmd := &cobra.Command{
		Use:   "templates",
		Short: "List available markdown templates",
		Run: func(cmd *cobra.Command, args []string) {
			availableTemplates := markdown.GetAvailableTemplates()

			fmt.Printf("Found %d templates available:\n", len(availableTemplates))
			for _, at := range availableTemplates {
				color.Green("  - %s", at)
			}
		},
	}

	return templatesCmd
}

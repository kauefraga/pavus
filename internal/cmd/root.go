package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/kauefraga/pavus/internal/markdown"
	"github.com/kauefraga/pavus/internal/server"
	"github.com/spf13/cobra"
)

func getRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "pavus",
		Short:   "The next-gen markdown tool",
		Version: "1.0.3",
		Example: "pavus\npavus path/to/markdown.md",
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			flagAssetDirectory, _ := cmd.Flags().GetString("asset-directory")

			mdPath := markdown.GetPath(args)
			assetDirectory := markdown.GetAssetDirectory(flagAssetDirectory)

			if mdPath == "" {
				color.Red("Error: there is no markdown to serve")
				os.Exit(1)
			}

			openBrowser, _ := cmd.Flags().GetBool("open-browser")

			err := server.ServeAndWatch(mdPath, assetDirectory, openBrowser)
			if err != nil {
				color.Red("Error: %s", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().StringP("asset-directory", "a", "", "assets directory to be served")
	rootCmd.Flags().BoolP("open-browser", "o", false, "open preview in the browser")

	return rootCmd
}

func Execute() {
	rootCmd := getRootCmd()

	if err := rootCmd.Execute(); err != nil {
		color.Red("Error: %s", err)
		os.Exit(1)
	}
}

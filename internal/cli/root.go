package cli

import (
	"fmt"
	"os"

	"github.com/kauefraga/pavus/internal/lib"
	"github.com/kauefraga/pavus/internal/server"
	"github.com/spf13/cobra"
)

func getRootCmd() *cobra.Command {
	var assetDirectory string

	rootCmd := &cobra.Command{
		Use:     "pavus",
		Short:   "The next-gen markdown tool",
		Version: "1.0.3",
		Example: "pavus\npavus path/to/markdown.md",
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			mdPath := lib.GetMarkdownPath(args)
			assetDirectory := lib.GetAssetDirectory(assetDirectory)

			if mdPath == "" {
				fmt.Println("Error: there is no markdown to serve")
				cmd.Help()
				os.Exit(0)
			}

			server.ServeAndWatch(mdPath, assetDirectory)
		},
	}

	rootCmd.Flags().StringVarP(&assetDirectory, "asset-directory", "a", "", "Define assets directory to be served")

	return rootCmd
}

func Execute() {
	err := getRootCmd().Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

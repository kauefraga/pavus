package cmd

import (
	"fmt"
	"os"

	"github.com/kauefraga/pavus/internal/lib"
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

			mdPath := lib.GetMarkdownPath(args)
			assetDirectory := lib.GetAssetDirectory(flagAssetDirectory)

			if mdPath == "" {
				fmt.Println("Error: there is no markdown to serve")
				cmd.Help()
				os.Exit(0)
			}

			err := server.ServeAndWatch(mdPath, assetDirectory)
			if err != nil {
				fmt.Println("Error:", err)
				os.Exit(1)
			}
		},
	}

	rootCmd.Flags().StringP("asset-directory", "a", "", "define assets directory to be served")

	return rootCmd
}

func Execute() {
	rootCmd := getRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

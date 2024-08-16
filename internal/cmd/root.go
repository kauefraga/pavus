package cmd

import (
	"fmt"
	"os"

	"github.com/kauefraga/pavus/internal/lib"
	"github.com/kauefraga/pavus/internal/server"
	"github.com/spf13/cobra"
)

func getRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "pavus",
		Short:   "The next-gen markdown tool",
		Example: "pavus\npavus path/to/markdown.md",
		Args:    cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			/* options
			 * - serve the first markdown (current)
			 * - serve all found files
			 * - serve the ones with priority (like readmes)
			 */

			mdPath := lib.GetMarkdownPath(args)

			if mdPath == "" {
				fmt.Println("Error: there is no markdown to serve")
				cmd.Help()
				os.Exit(0)
			}

			md, err := os.ReadFile(mdPath)
			if err != nil {
				fmt.Println("Error: failed to read the markdown file")
				os.Exit(1)
			}

			server.ServeAndWatch(md)
		},
	}
}

func Execute() {
	err := getRootCmd().Execute()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

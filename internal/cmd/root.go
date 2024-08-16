package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/kauefraga/pavus/internal/server"
	"github.com/spf13/cobra"
)

// Verify if a file is a markdown one
func isMarkdown(filename string) bool {
	isMarkdown, err := regexp.MatchString(`^.*\.(md)$`, filename)
	if err != nil {
		log.Fatalln(err)
	}

	return isMarkdown
}

// Verify if there is a markdown file and return the path of the first one
func findFirstMarkdownFile() string {
	var markdownPath string

	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if isMarkdown(info.Name()) {
			markdownPath = path
			return filepath.SkipAll
		}

		return err
	})
	if err != nil {
		log.Fatalln(err)
	}

	return markdownPath
}

func getMarkdownPath(args []string) string {
	if len(args) > 0 {
		info, err := os.Stat(args[0])
		if err != nil {
			return findFirstMarkdownFile()
		}

		if isMarkdown(info.Name()) {
			return args[0]
		}
	}

	return findFirstMarkdownFile()
}

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

			mdPath := getMarkdownPath(args)

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

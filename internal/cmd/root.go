package cmd

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"

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

// Verify if there is a markdown file nearby (same level or one level depth)
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

func getRootCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "pavus",
		Short: "The next-gen markdown tool",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			/* options
			 * - serve the first markdown (current)
			 * - serve all found files
			 * - serve the ones with priority (like readmes)
			 */

			mdPath := findFirstMarkdownFile()

			if len(args) == 0 && mdPath == "" {
				cmd.Help()
				os.Exit(0)
			}

			// if there is no file, check for a file path in the args
			if mdPath == "" {
				info, err := os.Stat(args[0])
				if err != nil {
					fmt.Println("error: there is no markdown file to serve")
					os.Exit(1)
				}

				if !isMarkdown(info.Name()) {
					fmt.Println("error: the file is not a markdown")
					os.Exit(1)
				}

				mdPath = args[0]
			}

			md, err := os.ReadFile(mdPath)
			if err != nil {
				fmt.Println("error: failed to read the markdown file")
				os.Exit(1)
			}

			fmt.Println(string(md))
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

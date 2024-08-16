package lib

import (
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
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

// If the first argument is a valid markdown path, return it,
// otherwise find and return the first markdown file path
func GetMarkdownPath(args []string) string {
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

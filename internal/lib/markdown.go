package lib

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
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

func MdToHTML(md []byte) template.HTML {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return template.HTML(markdown.Render(doc, renderer))
}

func ReadMarkdown(mdPath string) []byte {
	md, err := os.ReadFile(mdPath)
	if err != nil {
		fmt.Println("Error: failed to read the markdown file")
		os.Exit(1)
	}

	return md
}

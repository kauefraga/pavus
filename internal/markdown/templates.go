package markdown

import (
	"embed"
	"strings"
)

//go:embed templates/*.md
var templates embed.FS

// Transform template name from "Nice Template" to "nice-template"
func normalizeTemplateName(t string) string {
	return strings.ReplaceAll(
		strings.ToLower(t),
		" ",
		"-",
	)
}

func GetTemplate(t string) ([]byte, error) {
	return templates.ReadFile("templates/" + normalizeTemplateName(t) + ".md")
}

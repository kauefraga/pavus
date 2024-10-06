package markdown

import (
	"embed"
	"os"
	"strings"

	"github.com/fatih/color"
)

//go:embed templates/*.md
var templates embed.FS

// Transform template name from "Nice Template" to "nice-template"
func normalizeTemplateName(t string) string {
	return strings.ToLower(strings.ReplaceAll(t, " ", "-"))
}

// Transform template name from "nice-template" to "Nice Template"
func denormalizeTemplateName(t string) string {
	templateWords := strings.Split(t, "-")
	capitalizedTemplateWords := make([]string, len(templateWords))

	for i, tw := range templateWords {
		capitalizedTemplateWords[i] = strings.ToUpper(string(tw[0])) + tw[1:]
	}

	return strings.Join(capitalizedTemplateWords, " ")
}

func GetTemplate(t string) ([]byte, error) {
	return templates.ReadFile("templates/" + normalizeTemplateName(t) + ".md")
}

func GetAvailableTemplates() []string {
	var availableTemplates []string

	entries, err := templates.ReadDir("templates")
	if err != nil {
		color.Red("Error: failed to read templates directory")
		os.Exit(1)
	}

	for _, entry := range entries {
		templateName, _ := strings.CutSuffix(entry.Name(), ".md")

		availableTemplates = append(
			availableTemplates,
			denormalizeTemplateName(templateName),
		)
	}

	return availableTemplates
}

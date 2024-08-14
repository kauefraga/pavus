package server

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gomarkdown/markdown/parser"
)

// TODO: shutdown gracefully
// TODO: add styles (css)
// TODO: watch and reflect markdown changes

//go:embed layout.html
var layoutTemplate embed.FS

type LayoutData struct {
	Content template.HTML
}

func mdToHTML(md []byte) template.HTML {
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs | parser.NoEmptyLineBeforeBlock
	p := parser.NewWithExtensions(extensions)
	doc := p.Parse(md)

	htmlFlags := html.CommonFlags | html.HrefTargetBlank
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)

	return template.HTML(markdown.Render(doc, renderer))
}

func ServeAndWatch(md []byte) {
	tmpl, err := template.ParseFS(layoutTemplate, "layout.html")
	if err != nil {
		fmt.Println("error: failed to parse the layout template file")
		os.Exit(1)
	}

	data := LayoutData{
		Content: mdToHTML(md),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println("error:", err)
		}
	})

	fmt.Println("Waiting for you at http://localhost:3000... :)")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("error: failed to serve the markdown")
		os.Exit(1)
	}
}

package server

import (
	"embed"
	"errors"
	"html/template"
	"net/http"
	"os"

	"github.com/fatih/color"
)

//go:embed static/layout.html
var layout embed.FS

//go:embed static/*
var assets embed.FS

func ServeAndWatch(mdPath, assetDirectory string) error {
	http.HandleFunc("/sse", newReloadEventHandler(mdPath))
	http.Handle("/static/", http.FileServer(http.FS(assets)))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(assetDirectory))))

	tmpl, err := template.ParseFS(layout, "static/layout.html")
	if err != nil {
		return errors.New("failed to parse the layout template file")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		md, err := os.ReadFile(mdPath)
		if err != nil {
			color.Red("Error: failed to read the markdown file\n%s", err)
			return
		}

		err = tmpl.Execute(w, template.HTML(md))
		if err != nil {
			color.Red("Error: %s", err)
			return
		}
	})

	color.Yellow("[pavus] listening on http://localhost:3000 :)")

	return http.ListenAndServe(":3000", nil)
}

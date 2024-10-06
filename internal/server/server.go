package server

import (
	"embed"
	"errors"
	"html/template"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/pkg/browser"
)

var (
	//go:embed static/layout.html
	layout embed.FS

	//go:embed static/icon.png
	icon []byte
)

func ServeAndWatch(mdPath, assetDirectory string, openBrowser bool) error {
	http.HandleFunc("/sse", newReloadEventHandler(mdPath))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(assetDirectory))))

	http.HandleFunc("/icon.png", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/png")
		w.WriteHeader(http.StatusOK)
		w.Write(icon)
	})

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

	if openBrowser {
		color.Yellow("[pavus] opening preview in the browser...")

		browser.Stdout = nil
		browser.Stderr = nil
		err := browser.OpenURL("http://localhost:3000")
		if err != nil {
			color.Red("Error: failed to open the preview in the browser")
		}
	}

	return http.ListenAndServe(":3000", nil)
}

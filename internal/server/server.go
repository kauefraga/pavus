package server

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/kauefraga/pavus/internal/lib"
)

//go:embed static/layout.html
var layout embed.FS

//go:embed static/*
var assets embed.FS

func ServeAndWatch(mdPath, assetDirectory string) {
	http.HandleFunc("/sse", newServerSentEventsHandler(mdPath))
	http.Handle("/static/", http.FileServer(http.FS(assets)))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(assetDirectory))))

	tmpl, err := template.ParseFS(layout, "static/layout.html")
	if err != nil {
		fmt.Println("Error: failed to parse the layout template file")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, template.HTML(string(lib.ReadMarkdown(mdPath))))
		if err != nil {
			fmt.Println("Error:", err)
		}
	})

	fmt.Println("[pavus] happy writing!")
	fmt.Println("[pavus] waiting for you at http://localhost:3000... :)")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

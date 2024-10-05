package server

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
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
		fmt.Println("Error: failed to parse the layout template file")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		md, err := os.ReadFile(mdPath)
		if err != nil {
			fmt.Println("Error: failed to read the markdown file")
			os.Exit(1)
		}

		err = tmpl.Execute(w, template.HTML(md))
		if err != nil {
			fmt.Println("Error:", err)
		}
	})

	fmt.Println("[pavus] happy writing!")
	fmt.Println("[pavus] waiting for you at http://localhost:3000... :)")

	return http.ListenAndServe(":3000", nil)
}

package server

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//go:embed layout.html
var layoutTemplate embed.FS

func ServeAndWatch(markdown string) {
	tmpl, err := template.ParseFS(layoutTemplate, "layout.html")
	if err != nil {
		fmt.Println("error: failed to parse the layout template file")
		os.Exit(1)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = tmpl.Execute(w, markdown)
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

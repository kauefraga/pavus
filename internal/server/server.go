package server

import (
	"embed"
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/kauefraga/pavus/internal/lib"
)

// TODO: shutdown gracefully
// TODO: add styles (css)
// TODO: watch and reflect markdown changes

//go:embed previewer/layout.html
var layout embed.FS

type LayoutData struct {
	Content template.HTML
}

func ServeAndWatch(mdPath string) {
	http.HandleFunc("/ws", newWebSocketHandler(mdPath))

	tmpl, err := template.ParseFS(layout, "previewer/layout.html")
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

		data := LayoutData{
			Content: lib.MdToHTML(md),
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			fmt.Println("Error:", err)
		}
	})

	fmt.Println("Waiting for you at http://localhost:3000... :)")
	err = http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

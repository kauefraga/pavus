package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/fsnotify/fsnotify"
)

type ReloadEventMessage struct {
	Content string `json:"content"`
}

func sendReloadEvent(w http.ResponseWriter, mdPath string) {
	md, err := os.ReadFile(mdPath)
	if err != nil {
		color.Red("Error: failed to read the markdown file")
		os.Exit(1)
	}

	message := ReloadEventMessage{
		Content: string(md),
	}

	encodedMessage, err := json.Marshal(message)
	if err != nil {
		color.Red("Error: %s", err)
	}

	eventName := "event: reload\n"
	payloadMessage := fmt.Sprintf("data: %s\n\n", string(encodedMessage))

	w.Write([]byte(eventName))
	w.Write([]byte(payloadMessage))
}

func newReloadEventHandler(mdPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			color.Red("Error: %s", err)
		}
		defer watcher.Close()

		done := make(chan bool)

		color.Yellow("[pavus] watching: %s", mdPath)

		reloadToFixHappened := false

		go func() {
			for {
				if !reloadToFixHappened {
					reloadToFixHappened = true

					sendReloadEvent(w, mdPath)
					w.(http.Flusher).Flush()
				}

				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}

					if event.Has(fsnotify.Write) {
						color.Green("[pavus] reloading due to changes...")
						sendReloadEvent(w, mdPath)
					}

				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					color.Red("Error: %s", err)
					os.Exit(1)
				}

				w.(http.Flusher).Flush()
			}
		}()

		err = watcher.Add(mdPath)
		if err != nil {
			color.Red("Error: %s", err)
			os.Exit(1)
		}

		<-done
	}
}

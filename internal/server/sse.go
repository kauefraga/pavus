package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/kauefraga/pavus/internal/lib"
)

type ReloadEventMessage struct {
	Content string `json:"content"`
}

func newServerSentEventsHandler(mdPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		watcher, err := fsnotify.NewWatcher()
		if err != nil {
			fmt.Println("Error:", err)
		}
		defer watcher.Close()

		done := make(chan bool)

		fmt.Println("[pavus] watching:", mdPath)

		go func() {
			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}

					if event.Has(fsnotify.Write) {
						fmt.Println("[pavus] reloading due to changes...")

						message := ReloadEventMessage{
							Content: string(lib.ReadMarkdown(mdPath)),
						}

						encodedMessage, err := json.Marshal(message)
						if err != nil {
							log.Println("Error:", err)
						}

						eventName := "event: reload\n"
						payloadMessage := fmt.Sprintf("data: %s\n\n", string(encodedMessage))

						w.Write([]byte(eventName))
						w.Write([]byte(payloadMessage))
					}

				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					fmt.Println("error:", err)
					os.Exit(1)
				}

				w.(http.Flusher).Flush()
			}
		}()

		err = watcher.Add(mdPath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		<-done
	}
}

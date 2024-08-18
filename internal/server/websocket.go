package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
	"github.com/kauefraga/pavus/internal/lib"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool)

type Message struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func newWebSocketHandler(mdPath string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		socket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error: failed to upgrade connection (websocket)")
			os.Exit(1)
		}
		defer socket.Close()

		clients[socket] = true

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

						message := Message{
							Type:    "reload",
							Content: string(lib.ReadMarkdown(mdPath)),
						}

						socket.WriteJSON(message)
					}

				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					fmt.Println("error:", err)
					os.Exit(1)
				}
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

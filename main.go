package main

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/realtime-chat-go-react/pkg/websocket"
)

// Define websocket Endepoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

}

func main() {
	fmt.Println("Distibuted Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

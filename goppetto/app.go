package goppetto

import (
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
)

func init() {
	wsm := WebSocketManager{make(map[net.Addr]*websocket.Conn)}
	http.HandleFunc("/", index)
	http.HandleFunc("/specs", specs)

	http.HandleFunc("/api", wsm.ConnectionHandler)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func Run() {
	log.Printf("Start Goppetto on http://0.0.0.0:9999.")
	log.Fatal(http.ListenAndServe(":9999", nil))
}

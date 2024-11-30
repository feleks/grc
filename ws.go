package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")

	reader(ws)
}

func reader(conn *websocket.Conn) {
	handler := newHandler()
	for {
		_, msgRaw, err := conn.ReadMessage()
		if err != nil {
			log.Printf("failed to read message: %s", err.Error())
			return
		}

		err = handler.handle(conn, msgRaw)
		if err != nil {
			log.Printf("failed to handler message: %s", err.Error())
		}
	}
}

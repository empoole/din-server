package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = ":8080"
	SERVER_TYPE = "tcp"
)


var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients []websocket.Conn

func main() {
	fmt.Println("Server Running...")

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if(err != nil) {
			fmt.Println("Error: ")
			panic(err)
		}

		clients = append(clients, *conn)

		for {
			// Read message from browser
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			// Print the message to the console
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			for _, client := range clients {
				// Write message back to browser
				if err = client.WriteMessage(msgType, msg); err != nil {
					return
				}
			}

		}
	})

	http.ListenAndServe(":8080", nil)
}
package chat

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
	"main.go/internal/storage"
	"main.go/internal/types"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	SendPreviousMessages(conn)

	clients = append(clients, conn)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		var message types.Message
		err = msgpack.Unmarshal(msg, &message)
		if err != nil {
			fmt.Println(err)
			continue
		}

		NotifyClients(&message)
		storage.AddMessage(&message)
	}
}

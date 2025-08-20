package chat

import (
	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
	"main.go/internal/storage"
	"main.go/internal/types"
)

var clients []*websocket.Conn

func NotifyClients(msg *types.Message) {
	packedMsg, err := msgpack.Marshal(msg)
	if err != nil {
		return
	}

	for _, client := range clients {
		client.WriteMessage(websocket.BinaryMessage, packedMsg)
	}
}

func SendPreviousMessages(newConn *websocket.Conn) {

	var packedMessages [][]byte
	for _, msg := range storage.GetPreviousMessages() {
		packedMsg, err := msgpack.Marshal(msg)
		if err != nil {
			return
		}

		packedMessages = append(packedMessages, packedMsg)
	}

	for _, packedMsg := range packedMessages {
		newConn.WriteMessage(websocket.BinaryMessage, packedMsg)
	}

}

package chat

import (
	"strings"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
	"main.go/internal/storage"
	"main.go/internal/types"
)

var clients = make(map[string][]*websocket.Conn) // room code[connections]

func Init() {
	clients["global"] = []*websocket.Conn{}
}

func SendFirstConnectionMessage(user *websocket.Conn) {
	info := types.UserInfo{
		A: strings.Split(user.RemoteAddr().String(), ":")[0],
		B: "global",
	}

	msg := types.Message{A: info}

	packedMsg, err := msgpack.Marshal(msg)
	if err != nil {
		return
	}

	user.WriteMessage(websocket.BinaryMessage, packedMsg)
}

func NotifyClients(msg *types.Message) {
	packedMsg, err := msgpack.Marshal(msg)
	if err != nil {
		return
	}

	for _, client := range clients[msg.A.B] {
		client.WriteMessage(websocket.BinaryMessage, packedMsg)
	}
}

func SendPreviousMessages(roomCode string, newConn *websocket.Conn) {

	var packedMessages [][]byte
	for _, msg := range storage.GetPreviousMessages(roomCode) {
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

func ChangeRoom(user *websocket.Conn, oldRoom string, newRoom string) {
	if _, exists := clients[newRoom]; !exists {
		clients[newRoom] = []*websocket.Conn{}
	}

	for i, client := range clients[oldRoom] {
		if client == user {
			clients[oldRoom] = append(clients[oldRoom][:i], clients[oldRoom][i+1:]...)
			break
		}
	}

	clients[newRoom] = append(clients[newRoom], user)

	// Send previous messages for the new room
	SendPreviousMessages(newRoom, user)
}

package storage

import "main.go/internal/types"

var previousMessages = make(map[string][]*types.Message) // room code

func Init() {
	// Initialize the default "global" room
	previousMessages["global"] = []*types.Message{}
}

func GetPreviousMessages(roomCode string) []*types.Message {
	return previousMessages[roomCode]
}

func AddMessage(roomCode string, msg *types.Message) {
	previousMessages[roomCode] = append(previousMessages[roomCode], msg)
}

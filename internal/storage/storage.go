package storage

import "main.go/internal/types"

var previousMessages []*types.Message

func GetPreviousMessages() []*types.Message {
	return previousMessages
}

func AddMessage(msg *types.Message) {
	previousMessages = append(previousMessages, msg)
}

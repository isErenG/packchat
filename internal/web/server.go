package web

import (
	"log"
	"net/http"

	"main.go/internal/chat"
)

func SetupServer() {
	http.HandleFunc("/", GetHome)
	http.HandleFunc("/ws", chat.HandleWs)

	if err := http.ListenAndServe(":80", nil); err != nil {
		log.Fatal(err)
	}

}

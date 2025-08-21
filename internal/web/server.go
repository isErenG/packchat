package web

import (
	"log"
	"net/http"

	"main.go/internal/chat"
	"main.go/internal/storage"
)

func SetupServer() {

	chat.Init()
	storage.Init()

	// Set up static file server for all resources
	resourceFS := http.FileServer(http.Dir("../resources"))

	// Handle static files - CSS, JS, images, etc.
	http.Handle("/static/", http.StripPrefix("/static/", resourceFS))

	// Serve the home page
	http.HandleFunc("/", GetHome)
	http.HandleFunc("/ws", chat.HandleWs)

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

package web

import (
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../resources/home.html")
}

func GetChatRoom(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("chat room placeholder"))
}

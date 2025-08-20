package web

import (
	"net/http"
)

func GetHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "../resources/home.html")
}

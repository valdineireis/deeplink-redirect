package main

import (
	"net/http"
)

func handleDeeplink(w http.ResponseWriter, r *http.Request) {
	deeplink := r.URL.Query().Get("deeplink")
	if deeplink == "" {
		http.Error(w, "Missing deeplink parameter", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, deeplink, http.StatusFound)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/redirect", handleDeeplink)
	http.ListenAndServe(":8080", nil)
}

package main

import "net/http"

func handleDeeplink(w http.ResponseWriter, r *http.Request) {
	deeplink := r.URL.Query().Get("deeplink")
	if deeplink == "" {
		http.Error(w, "Missing deeplink parameter", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, deeplink, http.StatusFound)
}

func main() {
	http.HandleFunc("/redirect", handleDeeplink)
	http.ListenAndServe(":8080", nil)
}

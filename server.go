package main

import (
	"net/http"
	"os"
)

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/status", healthCheck)

	http.HandleFunc("/", redirect)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3060"
	}
	http.ListenAndServe(":"+port, nil)
}

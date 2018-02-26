package main

import (
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", redirect)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3060"
	}
	http.ListenAndServe(":"+port, nil)
}

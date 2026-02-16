package main

import (
	"log"
	"net/http"

	"hockey-game/internal/network"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("OK")); err != nil {
			log.Println("Health check write error:", err)
		}
	})

	http.HandleFunc("/ws", network.HandleWebSocket)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

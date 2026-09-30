package main

import (
	"log"
	"net/http"

	"github.com/nimisean/intelligent-systems-lab/backend/handlers"
)

func main() {
	http.HandleFunc("/health", handlers.Health)

	log.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

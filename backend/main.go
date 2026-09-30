package main

import (
	"log"
	"net/http"

	"github.com/nimisean/intelligent-systems-lab/backend/router"
)

func main() {
	r := router.New()

	log.Println("Server running on http://localhost:8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

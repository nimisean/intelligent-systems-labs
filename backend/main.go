package main

import (
	"log"
	"net/http"

	"github.com/YOUR-GITHUB-USERNAME/intelligent-systems-lab/backend/config"
	"github.com/YOUR-GITHUB-USERNAME/intelligent-systems-lab/backend/router"
)

func main() {
	cfg := config.Load()
	r := router.New()

	log.Printf("Server running on port %s", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal(err)
	}
}

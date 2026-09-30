package router

import (
	"net/http"

	"github.com/nimisean/intelligent-systems-lab/backend/handlers"
	"github.com/nimisean/intelligent-systems-lab/backend/middleware"
)

func New() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)

	return middleware.Logging(mux)
}

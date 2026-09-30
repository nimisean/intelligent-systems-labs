package router

import (
	"net/http"

	"github.com/nimisean/intelligent-systems-lab/backend/handlers"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handlers.Health)

	return mux
}

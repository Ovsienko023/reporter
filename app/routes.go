package app

import (
	"github.com/Ovsienko023/reporter/app/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetupRoutes(r chi.Router) http.Handler {
	r.Get("/reports", handlers.GetReport)

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})
	return r
}

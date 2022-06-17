package app

import (
	"github.com/Ovsienko023/reporter/app/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetupRoutes(r chi.Router) http.Handler {
	r.Get("/reports/{report_id}", handlers.GetReport)
	r.Post("/reports", handlers.CreateReport)
	return r
}

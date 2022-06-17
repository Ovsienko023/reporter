package app

import (
	"github.com/Ovsienko023/reporter/app/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func SetupRoutes(r chi.Router) http.Handler {
	r.Get("/api/v1/reports", handlers.GetReports)
	r.Get("/api/v1/reports/{report_id}", handlers.GetReport)
	r.Post("/api/v1/reports", handlers.CreateReport)
	return r
}

package http

import (
	"github.com/Ovsienko023/reporter/internal/report"

	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterHTTPEndpoints(router chi.Router, c report.Core) http.Handler {
	h := NewHandler(c)

	router.Get("/api/v1/reports", h.GetReports)
	router.Get("/api/v1/reports/{report_id}", h.GetReport)
	router.Post("/api/v1/reports", h.CreateReport)
	router.Put("/api/v1/reports/{report_id}", h.UpdateReport)

	return router
}

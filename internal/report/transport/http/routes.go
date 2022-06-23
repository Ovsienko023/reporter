package http

import (
	"github.com/Ovsienko023/reporter/internal/report"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterHTTPEndpoints(router chi.Router, core report.Core) http.Handler {
	h := NewHandler(core)

	router.Get("/api/v1/reports", h.GetReports)
	router.Get("/api/v1/reports/{report_id}", h.GetReport)
	router.Post("/api/v1/reports", h.CreateReport)

	return router
}

//func SetupRoutes(r chi.Router) http.Handler {
//	//r.Get("/api/v1/reports", handlers.GetReports)
//	//r.Get("/api/v1/reports/{report_id}", handlers.GetReport)
//	//r.Post("/api/v1/reports", handlers.CreateReport)
//	//r.Get("/api/v1/reports", .GetReports)
//	r.Get("/api/v1/reports/{report_id}", handlers.GetReport)
//	r.Post("/api/v1/reports", handlers.CreateReport)
//	return r
//}

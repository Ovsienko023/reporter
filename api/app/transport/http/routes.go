package http

import (
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RegisterHTTPEndpoints(router chi.Router, c core.Core) http.Handler {
	h := NewTransport(c)

	router.Post("/api/v1/login", h.GetToken)
	router.Post("/api/v1/sign_up", h.SignUp)

	router.Get("/api/v1/reports", h.GetReports)
	router.Get("/api/v1/reports/{report_id}", h.GetReport)
	router.Post("/api/v1/reports", h.CreateReport)
	router.Put("/api/v1/reports/{report_id}", h.UpdateReport)
	router.Delete("/api/v1/reports/{report_id}", h.DeleteReport)

	return router
}

package http

import (
	"fmt"
	"github.com/Ovsienko023/reporter/app/core"
	_ "github.com/Ovsienko023/reporter/docs"
	"github.com/Ovsienko023/reporter/infrastructure/configuration"
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func RegisterHTTPEndpoints(router chi.Router, c core.Core, apiConfig *configuration.Api) http.Handler {
	swaggerUrl := fmt.Sprintf("http://%s:%s/docs/doc.json", apiConfig.Doc.Host, apiConfig.Doc.Port)
	h := NewTransport(c)

	router.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(swaggerUrl)))

	router.Post("/api/v1/sign_in", h.SignIn)
	router.Post("/api/v1/sign_up", h.SignUp)

	router.Get("/api/v1/profile", h.GetProfile)
	router.Put("/api/v1/profile", h.UpdateProfile)

	router.Get("/api/v1/users", h.GetUsers)

	router.Get("/api/v1/reports", h.GetReports)
	router.Get("/api/v1/reports/{report_id}", h.GetReport)
	router.Post("/api/v1/reports", h.CreateReport)
	router.Put("/api/v1/reports/{report_id}", h.UpdateReport)
	router.Delete("/api/v1/reports/{report_id}", h.DeleteReport)
	router.Get("/api/v1/export/reports", h.ExportReportsToCsv)

	router.Get("/api/v1/stats", h.GetStatistics)

	router.Post("/api/v1/send_email", h.SendEmail)

	router.Post("/api/v1/users/{user_id}/permission", h.AddObjectToUserPermission)

	return router
}

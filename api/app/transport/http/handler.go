package http

import (
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/transport/http/handlers"
	"net/http"
)

type Transport struct {
	core core.Core
}

func NewTransport(c core.Core) *Transport {
	return &Transport{
		core: c,
	}
}

// AUTH

// SignIn ...  Sign In
// @Summary Sign In
// @Description Getting an authorization token
// @Tags Auth
// @Param request body domain.SignInRequest true "body params"
// @Success 200 {object} domain.SignInResponse
// @Success 401
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/sign_in [post]
func (t *Transport) SignIn(w http.ResponseWriter, r *http.Request) {
	handlers.SignIn(&t.core, w, r)
}

// SignUp ...  SignUp
// @Summary Sign Up
// @Description User registration
// @Tags Auth
// @Param request body domain.SignUpRequest true "body params"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/sign_up [post]
func (t *Transport) SignUp(w http.ResponseWriter, r *http.Request) {
	handlers.SignUp(&t.core, w, r)
}

// PROFILE

// GetProfile ...  Get Profile
// @Summary Get Profile
// @Description Getting user data
// @Tags Profile
// @Param request body domain.GetProfileRequest true "query params"
// @Success 200 {object} domain.GetProfileResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/profile [get]
func (t *Transport) GetProfile(w http.ResponseWriter, r *http.Request) {
	handlers.GetProfile(&t.core, w, r)
}

// UpdateProfile ...  Update Profile
// @Summary Update Profile
// @Description Updating user data
// @Tags Profile
// @Param request body domain.UpdateProfileRequest true "body params"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/profile [put]
func (t *Transport) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateProfile(&t.core, w, r)
}

// REPORTS

// GetReport ... Get report
// @Summary Get report
// @Description get report
// @Tags Reports
// @Param   report_id   path      string  true  "report_id"
// @Success 200 {object} domain.GetReportResponse
// @Failure 404 {object} httperror.ErrorResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/reports/{report_id} [get]
func (t *Transport) GetReport(w http.ResponseWriter, r *http.Request) {
	handlers.GetReport(&t.core, w, r)
}

// GetReports ... Get all reports
// @Summary Get all reports
// @Description get all reports
// @Tags Reports
// @Param request body domain.GetReportsRequest true "query params"
// @Success 200 {object} domain.GetReportsResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/reports [get]
func (t *Transport) GetReports(w http.ResponseWriter, r *http.Request) {
	handlers.GetReports(&t.core, w, r)
}

// CreateReport ...  Create report
// @Summary Create report
// @Description Create report
// @Tags Reports
// @Param request body domain.CreateReportRequest true "body params"
// @Success 201 {object} domain.CreateReportResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/reports [post]
func (t *Transport) CreateReport(w http.ResponseWriter, r *http.Request) {
	handlers.CreateReport(&t.core, w, r)
}

// UpdateReport ...  Update report
// @Summary Update report
// @Description Update report
// @Tags Reports
// @Param   id   path      string  true  "report_id"
// @Param request body domain.UpdateReportRequest true "body params"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/reports/{report_id} [put]
func (t *Transport) UpdateReport(w http.ResponseWriter, r *http.Request) {
	handlers.UpdateReport(&t.core, w, r)
}

// DeleteReport ...  Delete report
// @Summary Delete report
// @Description Delete report
// @Tags Reports
// @Param   id   path      string  true  "report_id"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/reports/{report_id} [delete]
func (t *Transport) DeleteReport(w http.ResponseWriter, r *http.Request) {
	handlers.DeleteReport(&t.core, w, r)
}

// ExportReportsToCsv ... Get all reports
// @Summary Export reports to csv
// @Description Export reports to csv
// @Tags Reports
// @Param request body domain.GetReportsRequest true "query params"
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/export/reports [get]
func (t *Transport) ExportReportsToCsv(w http.ResponseWriter, r *http.Request) {
	handlers.ExportReportsToCsv(&t.core, w, r)
}

// SendEmail ...  Sending a report by mail
// @Summary Send email
// @Description Send email
// @Tags Email
// @Param request body domain.SendEmailRequest true "body params"
// @Success 204
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/send_mail [post]
func (t *Transport) SendEmail(w http.ResponseWriter, r *http.Request) {
	handlers.SendEmail(&t.core, w, r)
}

// GetStatistics ...  Get Statistics
// @Summary Get Statistics
// @Description Get Statistics
// @Tags Statistics
// @Param request body domain.GetStatisticsRequest true "query params"
// @Success 200 {object} domain.GetStatisticsResponse
// @Failure 500 {object} httperror.ErrorResponse
// @Router /api/v1/stats [get]
func (t *Transport) GetStatistics(w http.ResponseWriter, r *http.Request) {
	handlers.GetStatistics(&t.core, w, r)
}

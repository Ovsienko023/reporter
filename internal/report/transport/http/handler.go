package http

import (
	"encoding/json"
	"errors"
	"github.com/Ovsienko023/reporter/internal/httperror"
	"github.com/Ovsienko023/reporter/internal/report"
	"github.com/Ovsienko023/reporter/internal/report/core"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	core report.Core
}

func NewHandler(c report.Core) *Handler {
	return &Handler{
		core: c,
	}
}

func (h *Handler) GetReport(w http.ResponseWriter, r *http.Request) {
	errorContainer := httperror.ErrorResponse{}

	ctx := r.Context()

	message := report.GetReportRequest{
		ReportId: chi.URLParam(r, "report_id"),
	}

	result, err := h.core.GetReport(ctx, &message)

	w.Header().Add("Content-Type", "application/json")

	if err != nil {
		switch {
		case errors.Is(err, core.ErrReportIdNotFound):
			errorContainer.Done(w, http.StatusNotFound, "report id not found")
			return
		}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

func (h *Handler) GetReports(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	message := report.GetReportsRequest{}
	result, _ := h.core.GetReports(ctx, &message)

	w.Header().Add("Content-Type", "application/json")

	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

func (h *Handler) CreateReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	var message report.CreateReportRequest

	err := decoder.Decode(&message)
	if err != nil {
		panic(err) //todo new httperror
	}

	w.Header().Add("Content-Type", "application/json")

	result, _ := h.core.CreateReport(ctx, &message)

	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

func (h *Handler) UpdateReport(w http.ResponseWriter, r *http.Request) {
	errorContainer := httperror.ErrorResponse{}

	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	var message report.UpdateReportRequest

	err := decoder.Decode(&message)
	if err != nil {
		panic(err) //todo new httperror
	}

	message.ReportId = chi.URLParam(r, "report_id")

	w.Header().Add("Content-Type", "application/json")

	err = h.core.UpdateReport(ctx, &message)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrReportIdNotFound):
			errorContainer.Done(w, http.StatusNotFound, "report id not found")
			return
		}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

}

func (h *Handler) DeleteReport(w http.ResponseWriter, r *http.Request) {
	errorContainer := httperror.ErrorResponse{}

	ctx := r.Context()

	message := report.DeleteReportRequest{
		ReportId: chi.URLParam(r, "report_id"),
	}

	message.ReportId = chi.URLParam(r, "report_id")

	w.Header().Add("Content-Type", "application/json")

	err := h.core.DeleteReport(ctx, &message)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrReportIdNotFound):
			errorContainer.Done(w, http.StatusNotFound, "report id not found")
			return
		}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}
}

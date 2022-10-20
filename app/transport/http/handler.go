package http

import (
	"encoding/json"
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	core core.Core
}

func NewHandler(c core.Core) *Handler {
	return &Handler{
		core: c,
	}
}

func (h *Handler) GetReport(w http.ResponseWriter, r *http.Request) {
	errorContainer := ErrorResponse{}
	w.Header().Add("Content-Type", "application/json")

	invokerId, err := authorize(r.Header.Get("Authorization"))
	if err != nil {
		errorContainer.Done(w, http.StatusUnauthorized, err.Error())
		return
	}

	message := domain.GetReportRequest{
		ReportId: chi.URLParam(r, "report_id"),
	}

	message.InvokerId = invokerId
	result, err := h.core.GetReport(r.Context(), &message)

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
	errorContainer := ErrorResponse{}
	w.Header().Add("Content-Type", "application/json")

	invokerId, err := authorize(r.Header.Get("Authorization"))
	if err != nil {
		errorContainer.Done(w, http.StatusUnauthorized, err.Error())
		return
	}

	message := domain.GetReportsRequest{
		InvokerId: invokerId,
	}
	result, _ := h.core.GetReports(r.Context(), &message)

	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

func (h *Handler) CreateReport(w http.ResponseWriter, r *http.Request) {
	errorContainer := ErrorResponse{}
	w.Header().Add("Content-Type", "application/json")

	invokerId, err := authorize(r.Header.Get("Authorization"))
	if err != nil {
		errorContainer.Done(w, http.StatusUnauthorized, err.Error())
		return
	}

	decoder := json.NewDecoder(r.Body)
	var message domain.CreateReportRequest

	err = decoder.Decode(&message)
	if err != nil {
		panic(err) //todo new httperror
	}

	message.InvokerId = invokerId
	result, _ := h.core.CreateReport(r.Context(), &message)

	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

func (h *Handler) UpdateReport(w http.ResponseWriter, r *http.Request) {
	errorContainer := ErrorResponse{}
	w.Header().Add("Content-Type", "application/json")

	invokerId, err := authorize(r.Header.Get("Authorization"))
	if err != nil {
		errorContainer.Done(w, http.StatusUnauthorized, err.Error())
		return
	}

	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	var message domain.UpdateReportRequest

	err = decoder.Decode(&message)
	if err != nil {
		panic(err) //todo new httperror
	}

	message.ReportId = chi.URLParam(r, "report_id")
	message.InvokerId = invokerId

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
	errorContainer := ErrorResponse{}
	w.Header().Add("Content-Type", "application/json")

	invokerId, err := authorize(r.Header.Get("Authorization"))
	if err != nil {
		errorContainer.Done(w, http.StatusUnauthorized, err.Error())
		return
	}

	message := domain.DeleteReportRequest{
		ReportId: chi.URLParam(r, "report_id"),
	}

	message.ReportId = chi.URLParam(r, "report_id")
	message.InvokerId = invokerId

	err = h.core.DeleteReport(r.Context(), &message)
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

func (h *Handler) GetToken(w http.ResponseWriter, r *http.Request) {
	errorContainer := ErrorResponse{}

	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	var message domain.GetTokenRequest

	err := decoder.Decode(&message)
	if err != nil {
		panic(err) //todo new httperror
	}

	w.Header().Add("Content-Type", "application/json")

	result, err := h.core.GetToken(ctx, &message)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrCredentials):
			errorContainer.Done(w, http.StatusForbidden, "permission denied")
			return
		}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}
	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

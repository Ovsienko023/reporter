package handlers

import (
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/auth"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func DeleteReport(c *core.Core, w http.ResponseWriter, r *http.Request) {
	errorContainer := httperror.ErrorResponse{}
	w.Header().Add("Content-Type", "application/json")

	invokerId, err := auth.Authorize(r.Header.Get("Authorization"))
	if err != nil {
		errorContainer.Done(w, http.StatusUnauthorized, err.Error())
		return
	}

	message := domain.DeleteReportRequest{
		ReportId: chi.URLParam(r, "report_id"),
	}

	message.ReportId = chi.URLParam(r, "report_id")
	message.InvokerId = invokerId

	err = c.DeleteReport(r.Context(), &message)
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

package handlers

import (
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"net/http"
	"strconv"
	"time"
)

func ExportReportsToCsv(c *core.Core, w http.ResponseWriter, r *http.Request) {
	errorContainer := httperror.ErrorResponse{}

	query := r.URL.Query()

	message := domain.ExportReportsRequest{
		Token: r.Header.Get("Authorization"),
	}

	// todo Вынести проверки на core + добавить http.Message (продумать генерацию docs)

	dateFrom := query.Get("date_from")
	if dateFrom != "" {
		i, err := strconv.ParseInt(dateFrom, 10, 64)
		if err != nil {
			errorContainer.Done(w, http.StatusBadRequest, "Invalid requests")
			return
		}
		tm := time.Unix(i, 0)
		message.DateFrom = &tm
	}

	dateTo := query.Get("date_to")
	if dateTo != "" {
		i, err := strconv.ParseInt(dateTo, 10, 64)
		if err != nil {
			errorContainer.Done(w, http.StatusBadRequest, "Invalid requests")
			return
		}
		tm := time.Unix(i, 0)
		message.DateTo = &tm
	}

	result, err := c.ExportReports(r.Context(), &message)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrUnauthorized):
			errorContainer.Done(w, http.StatusUnauthorized, err.Error())
			return
		}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

	err = FileResponse(w, result.Reports, "reports.csv")
	if err != nil {
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}
}

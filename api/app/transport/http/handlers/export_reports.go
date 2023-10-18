package handlers

import (
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"net/http"
)

func ExportReportsToCsv(c *core.Core, w http.ResponseWriter, r *http.Request) {
	errorContainer := httperror.ErrorResponse{}

	message := domain.ExportReportsToJsonRequest{
		Token: r.Header.Get("Authorization"),
	}

	result, err := c.ExportReportsToJson(r.Context(), &message)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrUnauthorized):
			errorContainer.Done(w, http.StatusUnauthorized, err.Error())
			return
		}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

	defer RemoveFile(result.File.Name())

	err = FileResponseWithReader(w, result.File, "export.json")
	if err != nil {
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}
}

package handlers

import (
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"net/http"
)

func GetStatistics(c *core.Core, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := httperror.ErrorResponse{}

	query := r.URL.Query()

	message := domain.ToGetStatisticsRequest(
		r.Header.Get("Authorization"),
		query.Get("from_date"),
		query.Get("to_date"),
	)

	result, err := c.GetStatistics(r.Context(), message)

	if err != nil {
		switch {
		case errors.Is(err, core.ErrUnauthorized):
			errorContainer.Done(w, http.StatusUnauthorized, err.Error())
			return
		}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

	ResponseMarshaller(w, http.StatusOK, result)
}

package handlers

import (
	"encoding/json"
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"net/http"
)

func CreateReport(c *core.Core, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := httperror.ErrorResponse{}

	decoder := json.NewDecoder(r.Body)
	var message domain.CreateReportRequest

	err := decoder.Decode(&message)
	if err != nil {
		errorContainer.Done(w, http.StatusInternalServerError, err.Error())
		return
	}

	message.Token = r.Header.Get("Authorization")
	result, err := c.CreateReport(r.Context(), &message)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrUnauthorized):
			errorContainer.Done(w, http.StatusUnauthorized, err.Error())
			return
		default:
			errorContainer.Done(w, http.StatusInternalServerError, err.Error())
			return
		}
	}

	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

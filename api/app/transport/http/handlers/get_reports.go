package handlers

import (
	"encoding/json"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/auth"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"net/http"
)

func GetReports(c *core.Core, w http.ResponseWriter, r *http.Request) {
	errorContainer := httperror.ErrorResponse{}
	w.Header().Add("Content-Type", "application/json")

	invokerId, err := auth.Authorize(r.Header.Get("Authorization"))
	if err != nil {
		errorContainer.Done(w, http.StatusUnauthorized, err.Error())
		return
	}

	message := domain.GetReportsRequest{
		InvokerId: invokerId,
	}
	result, _ := c.GetReports(r.Context(), &message)

	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

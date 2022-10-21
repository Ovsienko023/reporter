package handlers

import (
	"encoding/json"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/auth"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"net/http"
)

func CreateReport(c *core.Core, w http.ResponseWriter, r *http.Request) {
	errorContainer := httperror.ErrorResponse{}
	w.Header().Add("Content-Type", "application/json")

	invokerId, err := auth.Authorize(r.Header.Get("Authorization"))
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
	result, _ := c.CreateReport(r.Context(), &message)

	response, _ := json.Marshal(result)
	_, _ = w.Write(response)
}

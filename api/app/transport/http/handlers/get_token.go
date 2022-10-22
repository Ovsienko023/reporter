package handlers

import (
	"encoding/json"
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"net/http"
)

func GetToken(c *core.Core, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := httperror.ErrorResponse{}

	decoder := json.NewDecoder(r.Body)
	var message domain.GetTokenRequest

	err := decoder.Decode(&message)
	if err != nil {
		panic(err) //todo new httperror
	}

	result, err := c.GetToken(r.Context(), &message)
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

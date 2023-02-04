package handlers

import (
	"encoding/json"
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func RemoveObjectFromUserPermission(c *core.Core, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := httperror.ErrorResponse{}

	decoder := json.NewDecoder(r.Body)
	var message domain.RemoveObjectFromUserPermissionRequest

	err := decoder.Decode(&message)
	if err != nil {
		errorContainer.Done(w, http.StatusInternalServerError, err.Error())
		return
	}

	message.Token = r.Header.Get("Authorization")
	message.UserId = chi.URLParam(r, "user_id")

	err = c.RemoveObjectFromUserPermission(r.Context(), &message)
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

	JsonResponse(w, http.StatusNoContent, nil)
}
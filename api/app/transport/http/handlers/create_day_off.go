package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	validation "github.com/go-ozzo/ozzo-validation"
	"net/http"
)

func CreateDayOff(c *core.Core, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := httperror.ErrorResponse{}

	decoder := json.NewDecoder(r.Body)
	var message domain.CreateDayOffRequest

	err := decoder.Decode(&message)
	if err != nil {
		errorContainer.Done(w, http.StatusInternalServerError, err.Error())
		return
	}

	// todo Пренести валидацию на уровень core
	err = message.Validate()
	if s, ok := err.(validation.Errors); ok {
		for key, val := range s {
			errorContainer.Add(fmt.Sprintf("invalid.%s", key), val.Error(), key)
		}
		errorContainer.Done(w, http.StatusBadRequest, httperror.InvalidRequest)
		return
	}

	message.Token = r.Header.Get("Authorization")
	result, err := c.CreateDayOff(r.Context(), &message)
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

	JsonResponse(w, http.StatusCreated, result)
}

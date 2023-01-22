package handlers

import (
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
	"net/http"
	"strconv"
)

func GetUsers(c *core.Core, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := httperror.ErrorResponse{}

	query := r.URL.Query()

	message := domain.GetUsersRequest{
		Token: r.Header.Get("Authorization"),
	}

	// todo Вынести проверки на core + добавить http.Message (продумать генерацию docs)

	page := query.Get("page")
	if page != "" {
		intVar, err := strconv.Atoi(page)
		if err != nil {
			errorContainer.Done(w, http.StatusBadRequest, "Invalid requests")
			return
		}
		message.Page = ptr.Int(intVar)
	} else {
		message.Page = ptr.Int(1)
	}

	pageSize := query.Get("page_size")
	if pageSize != "" {
		intVar, err := strconv.Atoi(pageSize)
		if err != nil {
			errorContainer.Done(w, http.StatusBadRequest, "Invalid requests")
			return
		}
		message.PageSize = ptr.Int(intVar)
	} else {
		message.PageSize = ptr.Int(200)
	}

	result, err := c.GetUsers(r.Context(), &message)
	if err != nil {
		switch {
		case errors.Is(err, core.ErrUnauthorized):
			errorContainer.Done(w, http.StatusUnauthorized, err.Error())
			return
		}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

	JsonResponse(w, http.StatusOK, result)
}

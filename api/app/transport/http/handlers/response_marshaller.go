package handlers

import (
	"encoding/json"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"net/http"
)

func ResponseMarshaller(w http.ResponseWriter, code int, resp any) {
	if resp == nil {
		w.WriteHeader(code)
		return
	}

	response, err := json.Marshal(resp)
	if err != nil {
		errorContainer := httperror.ErrorResponse{}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.WriteHeader(code)
	_, _ = w.Write(response)
}

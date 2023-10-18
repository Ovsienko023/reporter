package handlers

import (
	"encoding/json"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"io"
	"net/http"
	"os"
)

func JsonResponse(w http.ResponseWriter, code int, resp any) {
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

func FileResponse(w http.ResponseWriter, file []byte, filename string) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)

	_, err := w.Write(file)
	if err != nil {
		return err
	}

	return nil
}

func FileResponseWithReader(w http.ResponseWriter, file io.Reader, filename string) error {
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)

	_, err := io.Copy(w, file)
	if err != nil {
		return err
	}

	return nil
}

func RemoveFile(fullPath string) error {
	if err := os.Remove(fullPath); err != nil {
		return err
	}

	return nil
}

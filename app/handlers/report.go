package handlers

import "net/http"

func GetReport(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("root."))
}

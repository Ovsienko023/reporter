package http

import (
	"encoding/json"
	"github.com/Ovsienko023/reporter/internal/httperror"
	"github.com/Ovsienko023/reporter/internal/report"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler struct {
	core report.Core
}

func NewHandler(c report.Core) *Handler {
	return &Handler{
		core: c,
	}
}

func (h *Handler) GetReport(w http.ResponseWriter, r *http.Request) {
	errorContainer := httperror.ErrorResponse{}

	ctx := r.Context()

	message := report.GetReportRequest{
		ReportId: chi.URLParam(r, "report_id"),
	}

	result, err := h.core.GetReport(ctx, &message) // todo add httperror
	if err != nil {
		errorContainer.Done(&w, http.StatusNotFound, "user not found")
		return
	}

	response, _ := json.Marshal(result)
	w.Write(response)
}

func (h *Handler) GetReports(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	message := report.GetReportsRequest{
		//InvokerId: "11111111-1111-1111-1111-111111111111", // todo
	}

	result, _ := h.core.GetReports(ctx, &message) // todo add httperror
	//reportsMSG := report.GetReportsResponse{
	//	Reports: []report.Report{}, //Reports: make([]report.Report, 0),
	//}
	//for _, obj := range result.Reports {
	//	reportsMSG.Reports = append(reportsMSG.Reports, obj)
	//}

	response, _ := json.Marshal(result)
	w.Write(response)
}

func (h *Handler) CreateReport(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	decoder := json.NewDecoder(r.Body)
	var message report.CreateReportRequest

	err := decoder.Decode(&message)
	if err != nil {
		panic(err) //todo new httperror
	}

	//message.InvokerId = "11111111-1111-1111-1111-111111111111" // todo
	result, _ := h.core.CreateReport(ctx, &message)
	//createdMSG := report.CreatedReportResponse{
	//	Id: result.Id,
	//}

	response, _ := json.Marshal(result)

	w.Write(response)
}

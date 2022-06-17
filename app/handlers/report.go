package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ReportTemplate struct {
	Variables map[string]string `json:"variables,omitempty"`
	Markup    string            `json:"markup,omitempty"`
}

type GetReportResponse struct {
	Id        string         `json:"id,omitempty"`
	CreatorId string         `json:"creator_id,omitempty"`
	CreatedAt int            `json:"created_at,omitempty"`
	UpdatedAt *int           `json:"updated_at,omitempty"`
	DeletedAt *int           `json:"deleted_at,omitempty"`
	StartTime int            `json:"start_time,omitempty"`
	EndTime   int            `json:"end_time,omitempty"`
	BreakTime int            `json:"break_time,omitempty"`
	WorkTime  int            `json:"work_time,omitempty"`
	Template  ReportTemplate `json:"template,omitempty"`
}

type CreateReportRequest struct {
	StartTime int            `json:"start_time,omitempty"`
	EndTime   int            `json:"end_time,omitempty"`
	BreakTime int            `json:"break_time,omitempty"`
	WorkTime  int            `json:"work_time,omitempty"`
	Template  ReportTemplate `json:"template,omitempty"`
}

type Store struct {
	Reports map[string]GetReportResponse
}

var store Store = Store{
	Reports: make(map[string]GetReportResponse),
}

func GetReport(w http.ResponseWriter, r *http.Request) {
	reportId := chi.URLParam(r, "report_id")

	if id, ok := store.Reports[reportId]; ok {
		response, _ := json.Marshal(id)
		w.Write(response)
	} else {
		w.Write([]byte("report_id not found"))
		return
	}

}

func CreateReport(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var message CreateReportRequest

	err := decoder.Decode(&message)
	if err != nil {
		panic(err)
	}

	id := uuid.New().String()

	store.Reports[id] = GetReportResponse{
		Id:        id,
		CreatorId: "0",
		CreatedAt: int(time.Now().Unix()),
		StartTime: message.StartTime,
		EndTime:   message.EndTime,
		BreakTime: message.BreakTime,
		WorkTime:  message.WorkTime,
		Template: ReportTemplate{
			Variables: message.Template.Variables,
			Markup:    message.Template.Markup,
		},
	}

	response, _ := json.Marshal(store.Reports[id])

	w.Write(response)
}

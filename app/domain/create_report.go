package domain

type CreateReportResponse struct {
	Id string `json:"id,omitempty"`
}

type CreateReportRequest struct {
	Title     string `json:"title,omitempty"`
	Date      int    `json:"date,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
	BreakTime int    `json:"break_time,omitempty"`
	WorkTime  int    `json:"work_time,omitempty"`
	Body      string `json:"body,omitempty"`
}

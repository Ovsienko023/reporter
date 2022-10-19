package domain

type CreateReportResponse struct {
	Id string `json:"id,omitempty"`
}

type CreateReportRequest struct {
	InvokerId string `json:"invoker_id,omitempty"`
	Title     string `json:"title,omitempty"`
	Date      int64  `json:"date,omitempty"`
	StartTime int64  `json:"start_time,omitempty"`
	EndTime   int64  `json:"end_time,omitempty"`
	BreakTime int64  `json:"break_time,omitempty"`
	WorkTime  int64  `json:"work_time,omitempty"`
	Body      string `json:"body,omitempty"`
}

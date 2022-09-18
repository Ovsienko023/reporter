package domain

type UpdateReportRequest struct {
	ReportId  string  `json:"id,omitempty"`
	Title     *string `json:"title,omitempty"`
	StartTime *int    `json:"start_time,omitempty"`
	EndTime   *int    `json:"end_time,omitempty"`
	BreakTime *int    `json:"break_time,omitempty"`
	WorkTime  *int    `json:"work_time,omitempty"`
	Body      *string `json:"body,omitempty"`
}

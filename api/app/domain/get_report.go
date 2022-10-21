package domain

type GetReportResponse struct {
	Report *Report `json:"report,omitempty"`
}

type Report struct {
	Id        *string `json:"id,omitempty"`
	Title     *string `json:"title,omitempty"`
	Date      *int64  `json:"date,omitempty"`
	CreatorId *string `json:"creator_id,omitempty"`
	CreatedAt *int64  `json:"created_at,omitempty"`
	UpdatedAt *int64  `json:"updated_at,omitempty"`
	StartTime *int64  `json:"start_time,omitempty"`
	EndTime   *int64  `json:"end_time,omitempty"`
	BreakTime *int64  `json:"break_time,omitempty"`
	WorkTime  *int64  `json:"work_time,omitempty"`
	Body      *string `json:"body,omitempty"`
}

type GetReportRequest struct {
	InvokerId string `json:"invoker_id,omitempty"`
	ReportId  string `json:"report_id,omitempty"`
}

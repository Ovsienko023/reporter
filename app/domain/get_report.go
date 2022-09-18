package domain

type GetReportResponse struct {
	Report *Report `json:"report,omitempty"`
}

type Report struct {
	Id        *string `json:"id,omitempty"`
	Title     *string `json:"title,omitempty"`
	CreatorId *string `json:"creator_id,omitempty"`
	CreatedAt *int    `json:"created_at,omitempty"`
	UpdatedAt *int    `json:"updated_at,omitempty"`
	DeletedAt *int    `json:"deleted_at,omitempty"`
	StartTime *int    `json:"start_time,omitempty"`
	EndTime   *int    `json:"end_time,omitempty"`
	BreakTime *int    `json:"break_time,omitempty"`
	WorkTime  *int    `json:"work_time,omitempty"`
	Body      *string `json:"body,omitempty"`
}

type GetReportRequest struct {
	ReportId string `json:"report_id,omitempty"`
}

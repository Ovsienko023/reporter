package domain

type GetReportsRequest struct{}

type GetReportsResponse struct {
	Count   *int         `json:"count,omitempty"`
	Reports []ReportItem `json:"reports" json:"reports,omitempty"`
}

type ReportItem struct {
	Id        *string `json:"id,omitempty"`
	Title     *string `json:"title,omitempty"`
	Date      *int64  `json:"date,omitempty"`
	CreatorId *string `json:"creator_id,omitempty"`
	CreatedAt *int64  `json:"created_at,omitempty"`
	UpdatedAt *int64  `json:"updated_at,omitempty"`
	DeletedAt *int64  `json:"deleted_at,omitempty"`
	StartTime *int64  `json:"start_time,omitempty"`
	EndTime   *int64  `json:"end_time,omitempty"`
	BreakTime *int64  `json:"break_time,omitempty"`
	WorkTime  *int64  `json:"work_time,omitempty"`
	Body      *string `json:"body,omitempty"`
}

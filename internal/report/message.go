package report

type GetReportResponse struct {
	Report Report `json:"report,omitempty"`
}

type GetReportsResponse struct {
	Reports []Report `json:"reports"`
}

type CreatedReportResponse struct {
	Id string `json:"id,omitempty"`
}

type GetReportRequest struct {
	ReportId string `json:"report_id,omitempty"`
}

type GetReport struct {
	InvokerId string `json:"invoker_id,omitempty"`
	ReportId  string `json:"report_id,omitempty"`
}

type GetReportsRequest struct {
}

type GetReports struct {
	InvokerId string `json:"invoker_id,omitempty"`
}

type CreateReportRequest struct {
	Title     string `json:"title,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
	BreakTime int    `json:"break_time,omitempty"`
	WorkTime  int    `json:"work_time,omitempty"`
	Body      string `json:"body,omitempty"`
}

type CreateReport struct {
	InvokerId string `json:"invoker_id,omitempty"`
	Title     string `json:"title,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
	BreakTime int    `json:"break_time,omitempty"`
	WorkTime  int    `json:"work_time,omitempty"`
	Body      string `json:"body,omitempty"`
}

type ReportTemplate struct {
	Variables map[string]string `json:"variables,omitempty"`
	Markup    string            `json:"markup,omitempty"`
}

type Report struct {
	Id        string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	CreatorId string `json:"creator_id,omitempty"`
	CreatedAt int    `json:"created_at,omitempty"`
	UpdatedAt *int   `json:"updated_at,omitempty"`
	DeletedAt *int   `json:"deleted_at,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
	BreakTime int    `json:"break_time,omitempty"`
	WorkTime  int    `json:"work_time,omitempty"`
	Body      string `json:"body,omitempty"`
}

type Reports struct {
	Reports []Report `json:"reports,omitempty"`
}

type CreatedReport struct {
	Id string `json:"id,omitempty"`
}

type UpdateReportRequest struct {
	ReportId  string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
	BreakTime int    `json:"break_time,omitempty"`
	WorkTime  int    `json:"work_time,omitempty"`
	Body      string `json:"body,omitempty"`
}

type UpdateReport struct {
	InvokerId string `json:"invoker_id,omitempty"`
	ReportId  string `json:"id,omitempty"`
	Title     string `json:"title,omitempty"`
	StartTime int    `json:"start_time,omitempty"`
	EndTime   int    `json:"end_time,omitempty"`
	BreakTime int    `json:"break_time,omitempty"`
	WorkTime  int    `json:"work_time,omitempty"`
	Body      string `json:"body,omitempty"`
}

type DeleteReportRequest struct {
	ReportId string `json:"report_id,omitempty"`
}

type DeleteReport struct {
	InvokerId string `json:"invoker_id,omitempty"`
	ReportId  string `json:"id,omitempty"`
}

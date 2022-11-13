package domain

import (
	"github.com/Ovsienko023/reporter/infrastructure/database"
	"time"
)

type UpdateReportRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	ReportId    string `json:"id,omitempty" swaggerignore:"true"`
	DisplayName string `json:"display_name,omitempty"`
	Date        int64  `json:"date,omitempty"`
	StartTime   int64  `json:"start_time,omitempty"`
	EndTime     int64  `json:"end_time,omitempty"`
	BreakTime   int64  `json:"break_time,omitempty"`
	WorkTime    int64  `json:"work_time,omitempty"`
	Body        string `json:"body,omitempty,"`
}

func (r *UpdateReportRequest) ToDbUpdateReport(invokerId string) *database.UpdateReport {
	return &database.UpdateReport{
		InvokerId:   invokerId,
		ReportId:    r.ReportId,
		DisplayName: r.DisplayName,
		Date:        time.Unix(r.Date, 0),
		StartTime:   time.Unix(r.StartTime, 0),
		EndTime:     time.Unix(r.EndTime, 0),
		BreakTime:   time.Unix(r.BreakTime, 0),
		WorkTime:    time.Unix(r.WorkTime, 0),
		Body:        r.Body,
	}
}

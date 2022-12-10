package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
	"time"
)

type UpdateReportRequest struct {
	Token       string  `json:"token,omitempty" swaggerignore:"true"`
	ReportId    string  `json:"id,omitempty" swaggerignore:"true"`
	DisplayName *string `json:"display_name,omitempty"`
	Date        *int64  `json:"date,omitempty"`
	StartTime   *int64  `json:"start_time,omitempty"`
	EndTime     *int64  `json:"end_time,omitempty"`
	BreakTime   *int64  `json:"break_time,omitempty"`
	WorkTime    *int64  `json:"work_time,omitempty"`
	Body        *string `json:"body,omitempty,"`
}

func (r *UpdateReportRequest) ToDbUpdateReport(invokerId string) *repository.UpdateReport {
	result := &repository.UpdateReport{
		InvokerId:   invokerId,
		ReportId:    r.ReportId,
		DisplayName: r.DisplayName,
		Body:        r.Body,
	}

	if r.Date != nil {
		result.Date = ptr.Time(time.Unix(*r.Date, 0).UTC())
	}
	if r.StartTime != nil {
		result.StartTime = ptr.Time(time.Unix(*r.StartTime, 0).UTC())
	}
	if r.EndTime != nil {
		result.EndTime = ptr.Time(time.Unix(*r.EndTime, 0).UTC())
	}
	if r.BreakTime != nil {
		result.BreakTime = ptr.Time(time.Unix(*r.BreakTime, 0).UTC())
	}
	if r.WorkTime != nil {
		result.WorkTime = ptr.Time(time.Unix(*r.WorkTime, 0).UTC())
	}

	return result
}

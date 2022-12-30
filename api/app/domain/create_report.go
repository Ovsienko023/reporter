package domain

import (
	"errors"
	"github.com/Ovsienko023/reporter/app/repository"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type CreateReportRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	DisplayName string `json:"display_name,omitempty"`
	Date        int64  `json:"date,omitempty"`
	StartTime   int64  `json:"start_time,omitempty"`
	EndTime     int64  `json:"end_time,omitempty"`
	BreakTime   int64  `json:"break_time,omitempty"`
	WorkTime    int64  `json:"work_time,omitempty"`
	Body        string `json:"body,omitempty"`
}

func (r CreateReportRequest) Validate() error {
	errs := validation.ValidateStruct(&r,
		// todo fields
		validation.Field(&r.StartTime, validation.Required),
		validation.Field(&r.EndTime, validation.Required),
	)

	if errs == nil && r.EndTime < r.StartTime {
		return errors.New("start_time must be less than end_time")
	}

	return errs
}

func (r *CreateReportRequest) ToDbCreateReport(invokerId string) *repository.CreateReport {
	date := time.Unix(r.Date, 0)
	return &repository.CreateReport{
		InvokerId:   invokerId,
		DisplayName: r.DisplayName,
		Date:        time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC),
		StartTime:   r.StartTime,
		EndTime:     r.EndTime,
		BreakTime:   r.BreakTime,
		WorkTime:    r.WorkTime,
		Body:        r.Body,
	}
}

type CreateReportResponse struct {
	Id string `json:"id,omitempty"`
}

func FromCreateReportResponse(resp *repository.CreatedReport) *CreateReportResponse {
	if resp == nil {
		return nil
	}

	return &CreateReportResponse{
		Id: resp.Id,
	}
}

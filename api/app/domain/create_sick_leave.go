package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type CreateSickLeaveRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	Date        int64  `json:"date,omitempty"`
	IsPaid      bool   `json:"is_paid,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r CreateSickLeaveRequest) Validate() error {
	errs := validation.ValidateStruct(&r) // todo fields
	//validation.Field(&r.StartTime, validation.Required),
	//validation.Field(&r.EndTime, validation.Required),

	return errs
}

func (r *CreateSickLeaveRequest) ToDbCreateSickLeave(invokerId string) *repository.CreateSickLeave {
	date := time.Unix(r.Date, 0)
	return &repository.CreateSickLeave{
		InvokerId:   invokerId,
		Date:        time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, time.UTC),
		IsPaid:      r.IsPaid,
		Status:      "approved", // todo move to const
		Description: r.Description,
	}
}

type CreateSickLeaveResponse struct {
	Id string `json:"id,omitempty"`
}

func FromCreateSickLeaveResponse(resp *repository.CreatedSickLeave) *CreateSickLeaveResponse {
	if resp == nil {
		return nil
	}

	return &CreateSickLeaveResponse{
		Id: resp.Id,
	}
}

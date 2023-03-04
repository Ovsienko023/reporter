package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type CreateVacationRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	DateFrom    int64  `json:"date_from,omitempty"`
	DateTo      int64  `json:"date_to,omitempty"`
	IsPaid      bool   `json:"is_paid,omitempty"`
	Description string `json:"description,omitempty"` // todo *
}

func (r CreateVacationRequest) Validate() error {
	errs := validation.ValidateStruct(&r,
		validation.Field(&r.DateFrom, validation.Required),
		validation.Field(&r.DateTo, validation.Required),
		validation.Field(&r.IsPaid, validation.Required),
	)
	return errs
}

func (r *CreateVacationRequest) ToDbCreateVacation(invokerId string) *repository.CreateVacation {
	dateFrom := time.Unix(r.DateFrom, 0)
	dateTo := time.Unix(r.DateTo, 0)

	return &repository.CreateVacation{
		InvokerId:   invokerId,
		DateFrom:    time.Date(dateFrom.Year(), dateFrom.Month(), dateFrom.Day(), 0, 0, 0, 0, time.UTC),
		DateTo:      time.Date(dateTo.Year(), dateTo.Month(), dateTo.Day(), 0, 0, 0, 0, time.UTC),
		IsPaid:      r.IsPaid,
		Status:      "approved", // todo move to const
		Description: r.Description,
	}
}

type CreateVacationResponse struct {
	Id string `json:"id,omitempty"`
}

func FromCreateVacationResponse(resp *repository.CreatedVacation) *CreateVacationResponse {
	if resp == nil {
		return nil
	}

	return &CreateVacationResponse{
		Id: resp.Id,
	}
}

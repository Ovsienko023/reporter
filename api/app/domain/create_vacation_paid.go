package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type CreateVacationPaidRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	DateFrom    int64  `json:"date_from,omitempty"`
	DateTo      int64  `json:"date_to,omitempty"`
	Description string `json:"description,omitempty"` // todo *
}

func (r CreateVacationPaidRequest) Validate() error {
	errs := validation.ValidateStruct(&r,
		validation.Field(&r.DateFrom, validation.Required),
		validation.Field(&r.DateTo, validation.Required),
	)
	return errs
}

func (r *CreateVacationPaidRequest) ToDbCreateVacationPaid(invokerId string) *repository.CreateVacationPaid {
	dateFrom := time.Unix(r.DateFrom, 0)
	dateTo := time.Unix(r.DateTo, 0)

	return &repository.CreateVacationPaid{
		InvokerId:   invokerId,
		DateFrom:    time.Date(dateFrom.Year(), dateFrom.Month(), dateFrom.Day(), 0, 0, 0, 0, time.UTC),
		DateTo:      time.Date(dateTo.Year(), dateTo.Month(), dateTo.Day(), 0, 0, 0, 0, time.UTC),
		Status:      "approved", // todo move to const
		Description: r.Description,
	}
}

type CreateVacationPaidResponse struct {
	Id string `json:"id,omitempty"`
}

func FromCreateVacationPaidResponse(resp *repository.CreatedVacationPaid) *CreateVacationPaidResponse {
	if resp == nil {
		return nil
	}

	return &CreateVacationPaidResponse{
		Id: resp.Id,
	}
}

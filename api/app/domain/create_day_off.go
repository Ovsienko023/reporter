package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	validation "github.com/go-ozzo/ozzo-validation"
	"time"
)

type CreateDayOffRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	DateFrom    int64  `json:"date_from,omitempty"`
	DateTo      int64  `json:"date_to,omitempty"`
	Description string `json:"description,omitempty"`
}

func (r CreateDayOffRequest) Validate() error {
	errs := validation.ValidateStruct(&r,
		validation.Field(&r.DateFrom, validation.Required),
		validation.Field(&r.DateTo, validation.Required),
	)

	return errs
}

func (r *CreateDayOffRequest) ToDbCreateDayOff(invokerId string) *repository.CreateDayOff {
	dateFrom := time.Unix(r.DateFrom, 0)
	dateTo := time.Unix(r.DateTo, 0)
	return &repository.CreateDayOff{
		InvokerId:   invokerId,
		DateFrom:    time.Date(dateFrom.Year(), dateFrom.Month(), dateFrom.Day(), 0, 0, 0, 0, time.UTC),
		DateTo:      time.Date(dateTo.Year(), dateTo.Month(), dateTo.Day(), 0, 0, 0, 0, time.UTC),
		Status:      "approved", // todo move to const
		Description: r.Description,
	}
}

type CreateDayOffResponse struct {
	Id string `json:"id,omitempty"`
}

func FromCreateDayOffResponse(resp *repository.CreatedDayOff) *CreateDayOffResponse {
	if resp == nil {
		return nil
	}

	return &CreateDayOffResponse{
		Id: resp.Id,
	}
}

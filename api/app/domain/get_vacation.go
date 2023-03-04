package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

type GetVacationRequest struct {
	Token      string `json:"token,omitempty" swaggerignore:"true"`
	VacationId string `json:"vacation_id,omitempty" swaggerignore:"true"`
}

func (r *GetVacationRequest) ToDbGetVacation(invokerId string) *repository.GetVacation {
	return &repository.GetVacation{
		InvokerId:  invokerId,
		VacationId: r.VacationId,
	}
}

type GetVacationResponse struct {
	Vacation *Vacation `json:"vacation,omitempty"`
}

type Vacation struct {
	Id          *string `json:"id,omitempty"`
	DateFrom    *int64  `json:"date_from,omitempty"`
	DateTo      *int64  `json:"date_to,omitempty"`
	IsPaid      *bool   `json:"is_paid,omitempty"`
	Status      *string `json:"status,omitempty"`
	CreatorId   *string `json:"creator_id,omitempty"`
	Description *string `json:"description,omitempty"`
}

func FromGetVacationResponse(resp *repository.Vacation) *GetVacationResponse {
	if resp == nil {
		return nil
	}

	return &GetVacationResponse{
		Vacation: &Vacation{
			Id:          resp.Id,
			DateFrom:    ptr.Int64(resp.DateFrom.Unix()),
			DateTo:      ptr.Int64(resp.DateTo.Unix()),
			CreatorId:   resp.CreatorId,
			IsPaid:      resp.IsPaid,
			Status:      resp.Status,
			Description: resp.Description,
		},
	}
}

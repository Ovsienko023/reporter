package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

type GetVacationRequest struct {
	Token      string `json:"token,omitempty" swaggerignore:"true"`
	UserId     string `json:"user_id,omitempty" swaggerignore:"true"`
	VacationId string `json:"vacation_id,omitempty" swaggerignore:"true"`
}

func (r *GetVacationRequest) ToDbGetVacation(invokerId string) *repository.GetVacation {
	return &repository.GetVacation{
		InvokerId:  invokerId,
		UserId:     r.UserId,
		VacationId: r.VacationId,
	}
}

type GetVacationResponse struct {
	Vacation *Vacation `json:"vacation,omitempty"`
}

type Vacation struct {
	Id          *string `json:"id,omitempty"`
	Date        *int64  `json:"date,omitempty"`
	IsPaid      *bool   `json:"is_paid,omitempty"`
	State       *string `json:"state,omitempty"`
	Status      *string `json:"status,omitempty"`
	Description *string `json:"description,omitempty"`
}

func FromGetVacationResponse(resp *repository.Vacation) *GetVacationResponse {
	if resp == nil {
		return nil
	}

	return &GetVacationResponse{
		Vacation: &Vacation{
			Id:          resp.Id,
			Date:        ptr.Int64(resp.Date.Unix()),
			IsPaid:      resp.IsPaid,
			State:       resp.State,
			Status:      resp.Status,
			Description: resp.Description,
		},
	}
}

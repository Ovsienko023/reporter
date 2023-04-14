package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

type GetVacationPaidRequest struct {
	Token          string `json:"token,omitempty" swaggerignore:"true"`
	VacationPaidId string `json:"vacation_paid_id,omitempty" swaggerignore:"true"`
}

func (r *GetVacationPaidRequest) ToDbGetVacationPaid(invokerId string) *repository.GetVacationPaid {
	return &repository.GetVacationPaid{
		InvokerId:      invokerId,
		VacationPaidId: r.VacationPaidId,
	}
}

type GetVacationPaidResponse struct {
	VacationPaid *VacationPaid `json:"vacation_paid,omitempty"`
}

type VacationPaid struct {
	Id          *string `json:"id,omitempty"`
	DateFrom    *int64  `json:"date_from,omitempty"`
	DateTo      *int64  `json:"date_to,omitempty"`
	Status      *string `json:"status,omitempty"`
	CreatorId   *string `json:"creator_id,omitempty"`
	Description *string `json:"description,omitempty"`
}

func FromGetVacationPaidResponse(resp *repository.VacationPaid) *GetVacationPaidResponse {
	if resp == nil {
		return nil
	}

	return &GetVacationPaidResponse{
		VacationPaid: &VacationPaid{
			Id:          resp.Id,
			DateFrom:    ptr.Int64(resp.DateFrom.Unix()),
			DateTo:      ptr.Int64(resp.DateTo.Unix()),
			CreatorId:   resp.CreatorId,
			Status:      resp.Status,
			Description: resp.Description,
		},
	}
}

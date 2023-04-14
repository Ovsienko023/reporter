package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

type GetVacationUnpaidRequest struct {
	Token            string `json:"token,omitempty" swaggerignore:"true"`
	VacationUnpaidId string `json:"vacation_unpaid_id,omitempty" swaggerignore:"true"`
}

func (r *GetVacationUnpaidRequest) ToDbGetVacationUnpaid(invokerId string) *repository.GetVacationUnpaid {
	return &repository.GetVacationUnpaid{
		InvokerId:        invokerId,
		VacationUnpaidId: r.VacationUnpaidId,
	}
}

type GetVacationUnpaidResponse struct {
	VacationUnpaid *VacationUnpaid `json:"vacation_unpaid,omitempty"`
}

type VacationUnpaid struct {
	Id          *string `json:"id,omitempty"`
	DateFrom    *int64  `json:"date_from,omitempty"`
	DateTo      *int64  `json:"date_to,omitempty"`
	Status      *string `json:"status,omitempty"`
	CreatorId   *string `json:"creator_id,omitempty"`
	Description *string `json:"description,omitempty"`
}

func FromGetVacationUnpaidResponse(resp *repository.VacationUnpaid) *GetVacationUnpaidResponse {
	if resp == nil {
		return nil
	}

	return &GetVacationUnpaidResponse{
		VacationUnpaid: &VacationUnpaid{
			Id:          resp.Id,
			DateFrom:    ptr.Int64(resp.DateFrom.Unix()),
			DateTo:      ptr.Int64(resp.DateTo.Unix()),
			CreatorId:   resp.CreatorId,
			Status:      resp.Status,
			Description: resp.Description,
		},
	}
}

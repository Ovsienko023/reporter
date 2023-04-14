package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
)

type DeleteVacationUnpaidRequest struct {
	Token            string `json:"token,omitempty"`
	VacationUnpaidId string `json:"vacation_unpaid_id,omitempty"`
}

func (r *DeleteVacationUnpaidRequest) ToDbDeleteVacationUnpaid(invokerId string) *repository.DeleteVacationUnpaid {
	return &repository.DeleteVacationUnpaid{
		InvokerId:        invokerId,
		VacationUnpaidId: r.VacationUnpaidId,
	}
}

package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
)

type DeleteVacationPaidRequest struct {
	Token          string `json:"token,omitempty"`
	VacationPaidId string `json:"vacation_paid_id,omitempty"`
}

func (r *DeleteVacationPaidRequest) ToDbDeleteVacationPaid(invokerId string) *repository.DeleteVacationPaid {
	return &repository.DeleteVacationPaid{
		InvokerId:      invokerId,
		VacationPaidId: r.VacationPaidId,
	}
}

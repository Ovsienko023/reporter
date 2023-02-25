package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
)

type DeleteVacationRequest struct {
	Token      string `json:"token,omitempty"`
	VacationId string `json:"vacation_id,omitempty"`
}

func (r *DeleteVacationRequest) ToDbDeleteVacation(invokerId string) *repository.DeleteVacation {
	return &repository.DeleteVacation{
		InvokerId:  invokerId,
		VacationId: r.VacationId,
	}
}

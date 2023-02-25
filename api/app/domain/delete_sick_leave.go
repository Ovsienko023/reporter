package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
)

type DeleteSickLeaveRequest struct {
	Token       string `json:"token,omitempty"`
	SickLeaveId string `json:"sick_leave_id,omitempty"`
}

func (r *DeleteSickLeaveRequest) ToDbDeleteSickLeave(invokerId string) *repository.DeleteSickLeave {
	return &repository.DeleteSickLeave{
		InvokerId:   invokerId,
		SickLeaveId: r.SickLeaveId,
	}
}

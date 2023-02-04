package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

type GetSickLeaveRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	UserId      string `json:"user_id,omitempty" swaggerignore:"true"`
	SickLeaveId string `json:"sick_leave_id,omitempty" swaggerignore:"true"`
}

func (r *GetSickLeaveRequest) ToDbGetSickLeave(invokerId string) *repository.GetSickLeave {
	return &repository.GetSickLeave{
		InvokerId:   invokerId,
		UserId:      r.UserId,
		SickLeaveId: r.SickLeaveId,
	}
}

type GetSickLeaveResponse struct {
	SickLeave *SickLeave `json:"sick_leave,omitempty"`
}

type SickLeave struct {
	Id          *string `json:"id,omitempty"`
	Date        *int64  `json:"date,omitempty"`
	IsPaid      *bool   `json:"is_paid,omitempty"`
	State       *string `json:"state,omitempty"`
	Status      *string `json:"status,omitempty"`
	Description *string `json:"description,omitempty"`
}

func FromGetSickLeaveResponse(resp *repository.SickLeave) *GetSickLeaveResponse {
	if resp == nil {
		return nil
	}

	return &GetSickLeaveResponse{
		SickLeave: &SickLeave{
			Id:          resp.Id,
			Date:        ptr.Int64(resp.Date.Unix()),
			IsPaid:      resp.IsPaid,
			State:       resp.State,
			Status:      resp.Status,
			Description: resp.Description,
		},
	}
}

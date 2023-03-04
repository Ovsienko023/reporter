package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

type GetSickLeaveRequest struct {
	Token       string `json:"token,omitempty" swaggerignore:"true"`
	SickLeaveId string `json:"sick_leave_id,omitempty" swaggerignore:"true"`
}

func (r *GetSickLeaveRequest) ToDbGetSickLeave(invokerId string) *repository.GetSickLeave {
	return &repository.GetSickLeave{
		InvokerId:   invokerId,
		SickLeaveId: r.SickLeaveId,
	}
}

type GetSickLeaveResponse struct {
	SickLeave *SickLeave `json:"sick_leave,omitempty"`
}

type SickLeave struct {
	Id          *string `json:"id,omitempty"`
	CreatorId   *string `json:"creator_id,omitempty"`
	DateFrom    *int64  `json:"date_from,omitempty"`
	DateTo      *int64  `json:"date_to,omitempty"`
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
			DateFrom:    ptr.Int64(resp.DateFrom.Unix()),
			DateTo:      ptr.Int64(resp.DateTo.Unix()),
			CreatorId:   resp.CreatorId,
			Status:      resp.Status,
			Description: resp.Description,
		},
	}
}

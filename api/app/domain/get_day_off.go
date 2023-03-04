package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

type GetDayOffRequest struct {
	Token    string `json:"token,omitempty" swaggerignore:"true"`
	DayOffId string `json:"day_off_id,omitempty" swaggerignore:"true"`
}

func (r *GetDayOffRequest) ToDbGetDayOff(invokerId string) *repository.GetDayOff {
	return &repository.GetDayOff{
		InvokerId: invokerId,
		DayOffId:  r.DayOffId,
	}
}

type GetDayOffResponse struct {
	DayOff *DayOff `json:"sick_leave,omitempty"`
}

type DayOff struct {
	Id          *string `json:"id,omitempty"`
	CreatorId   *string `json:"creator_id,omitempty"`
	DateFrom    *int64  `json:"date_from,omitempty"`
	DateTo      *int64  `json:"date_to,omitempty"`
	Status      *string `json:"status,omitempty"`
	Description *string `json:"description,omitempty"`
}

func FromGetDayOffResponse(resp *repository.DayOff) *GetDayOffResponse {
	if resp == nil {
		return nil
	}

	return &GetDayOffResponse{
		DayOff: &DayOff{
			Id:          resp.Id,
			DateFrom:    ptr.Int64(resp.DateFrom.Unix()),
			DateTo:      ptr.Int64(resp.DateTo.Unix()),
			CreatorId:   resp.CreatorId,
			Status:      resp.Status,
			Description: resp.Description,
		},
	}
}

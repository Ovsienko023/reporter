package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
)

type DeleteDayOffRequest struct {
	Token    string `json:"token,omitempty"`
	DayOffId string `json:"day_off_id,omitempty"`
}

func (r *DeleteDayOffRequest) ToDbDeleteDayOff(invokerId string) *repository.DeleteDayOff {
	return &repository.DeleteDayOff{
		InvokerId: invokerId,
		DayOffId:  r.DayOffId,
	}
}

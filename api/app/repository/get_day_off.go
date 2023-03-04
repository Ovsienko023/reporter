package repository

import (
	"context"
	"time"
)

const sqlGetDayOff = `
	select id,
		   date_from,
		   date_to,
		   status,
		   description,
		   creator_id,
		   created_at,
		   updated_at,
		   payload
    from main.day_off as sl
    where id = $1 and exists(select 1
        					 from main.permissions_users_to_objects as pto
            				 where pto.user_id = $2 and
       							pto.object_id =  sl.creator_id)`

func (c *Client) GetDayOff(ctx context.Context, msg *GetDayOff) (*DayOff, error) {
	isAuth, err := c.checkUserPermission(ctx, msg.InvokerId, msg.InvokerId)
	if !isAuth {
		return nil, ErrPermissionDenied
	}

	row, err := c.driver.Query(ctx,
		sqlGetDayOff,
		msg.DayOffId,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	dayOff := &DayOff{}

	for row.Next() {
		err := row.Scan(
			&dayOff.Id,
			&dayOff.DateFrom,
			&dayOff.DateTo,
			&dayOff.Status,
			&dayOff.Description,
			&dayOff.CreatorId,
			&dayOff.CreatedAt,
			&dayOff.UpdatedAt,
			&dayOff.Payload,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	if dayOff.Id == nil {
		return nil, ErrDayOffIdNotFound
	}

	return dayOff, nil
}

type GetDayOff struct {
	InvokerId string `json:"invoker_id,omitempty"`
	DayOffId  string `json:"day_off_id,omitempty"`
}

type DayOff struct {
	Id          *string        `json:"id,omitempty"`
	DateFrom    *time.Time     `json:"date_from,omitempty"`
	DateTo      *time.Time     `json:"date_to,omitempty"`
	Status      *string        `json:"status,omitempty"`
	Description *string        `json:"description,omitempty"`
	CreatorId   *string        `json:"creator_id,omitempty"`
	CreatedAt   *time.Time     `json:"created_at,omitempty"`
	UpdatedAt   *time.Time     `json:"updated_at,omitempty"`
	Payload     *DayOffPayload `json:"payload,omitempty"`
}

type DayOffPayload struct{}

package repository

import (
	"context"
	"time"
)

const sqlGetSickLeave = `
	select id,
		   date_from,
		   date_to,
		   status,
		   description,
		   creator_id,
		   created_at,
		   updated_at,
		   payload
    from main.sick_leave as sl
    where id = $1 and exists(select 1
        					 from main.permissions_users_to_objects as pto
            				 where pto.user_id = $2 and
       							pto.object_id =  sl.creator_id)`

func (c *Client) GetSickLeave(ctx context.Context, msg *GetSickLeave) (*SickLeave, error) {
	isAuth, err := c.checkUserPermission(ctx, msg.InvokerId, msg.InvokerId)
	if !isAuth {
		return nil, ErrPermissionDenied
	}

	row, err := c.driver.Query(ctx,
		sqlGetSickLeave,
		msg.SickLeaveId,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	sickLeave := &SickLeave{}

	for row.Next() {
		err := row.Scan(
			&sickLeave.Id,
			&sickLeave.DateFrom,
			&sickLeave.DateTo,
			&sickLeave.Status,
			&sickLeave.Description,
			&sickLeave.CreatorId,
			&sickLeave.CreatedAt,
			&sickLeave.UpdatedAt,
			&sickLeave.Payload,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	if sickLeave.Id == nil {
		return nil, ErrSickLeaveIdNotFound
	}

	return sickLeave, nil
}

type GetSickLeave struct {
	InvokerId   string `json:"invoker_id,omitempty"`
	SickLeaveId string `json:"sick_leave_id,omitempty"`
}

type SickLeave struct {
	Id          *string           `json:"id,omitempty"`
	DateFrom    *time.Time        `json:"date_from,omitempty"`
	DateTo      *time.Time        `json:"date_to,omitempty"`
	Status      *string           `json:"status,omitempty"`
	Description *string           `json:"description,omitempty"`
	CreatorId   *string           `json:"creator_id,omitempty"`
	CreatedAt   *time.Time        `json:"created_at,omitempty"`
	UpdatedAt   *time.Time        `json:"updated_at,omitempty"`
	Payload     *SickLeavePayload `json:"payload,omitempty"`
}

type SickLeavePayload struct{}

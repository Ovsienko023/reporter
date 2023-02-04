package repository

import (
	"context"
	"time"
)

const sqlGetSickLeave = `
	select id,
		   date,
		   is_paid,
		   state,
		   status,
		   description,
		   creator_id,
		   created_at,
		   updated_at,
		   payload
    from main.sick_leave
    inner join main.sick_leave_to_users stu on sick_leave.id = stu.sick_leave_id
    where id = $1 and 
          stu.user_id = $2`

func (c *Client) GetSickLeave(ctx context.Context, msg *GetSickLeave) (*SickLeave, error) {
	isAuth, err := c.checkUserPermission(ctx, msg.InvokerId, msg.UserId)
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
			&sickLeave.Date,
			&sickLeave.IsPaid,
			&sickLeave.State,
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
	UserId      string `json:"user_id,omitempty"`
	SickLeaveId string `json:"sick_leave_id,omitempty"`
}

type SickLeave struct {
	Id          *string           `json:"id,omitempty"`
	Date        *time.Time        `json:"date,omitempty"`
	IsPaid      *bool             `json:"is_paid,omitempty"`
	State       *string           `json:"state,omitempty"`
	Status      *string           `json:"status,omitempty"`
	Description *string           `json:"description,omitempty"`
	CreatorId   *string           `json:"creator_id,omitempty"`
	CreatedAt   *time.Time        `json:"created_at,omitempty"`
	UpdatedAt   *time.Time        `json:"updated_at,omitempty"`
	Payload     *SickLeavePayload `json:"payload,omitempty"`
}

type SickLeavePayload struct{}

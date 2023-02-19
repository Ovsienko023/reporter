package repository

import (
	"context"
	"time"
)

const sqlGetVacation = `
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
    from main.vacation
    inner join main.vacation_to_users stu on vacation.id = stu.vacation_id
    where id = $1 and 
          stu.user_id = $2`

func (c *Client) GetVacation(ctx context.Context, msg *GetVacation) (*Vacation, error) {
	isAuth, err := c.checkUserPermission(ctx, msg.InvokerId, msg.UserId)
	if !isAuth {
		return nil, ErrPermissionDenied
	}

	row, err := c.driver.Query(ctx,
		sqlGetVacation,
		msg.VacationId,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	vacation := &Vacation{}

	for row.Next() {
		err := row.Scan(
			&vacation.Id,
			&vacation.Date,
			&vacation.IsPaid,
			&vacation.State,
			&vacation.Status,
			&vacation.Description,
			&vacation.CreatorId,
			&vacation.CreatedAt,
			&vacation.UpdatedAt,
			&vacation.Payload,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	if vacation.Id == nil {
		return nil, ErrVacationIdNotFound
	}

	return vacation, nil
}

type GetVacation struct {
	InvokerId  string `json:"invoker_id,omitempty"`
	UserId     string `json:"user_id,omitempty"`
	VacationId string `json:"vacation_id,omitempty"`
}

type Vacation struct {
	Id          *string          `json:"id,omitempty"`
	Date        *time.Time       `json:"date,omitempty"`
	IsPaid      *bool            `json:"is_paid,omitempty"`
	State       *string          `json:"state,omitempty"`
	Status      *string          `json:"status,omitempty"`
	Description *string          `json:"description,omitempty"`
	CreatorId   *string          `json:"creator_id,omitempty"`
	CreatedAt   *time.Time       `json:"created_at,omitempty"`
	UpdatedAt   *time.Time       `json:"updated_at,omitempty"`
	Payload     *VacationPayload `json:"payload,omitempty"`
}

type VacationPayload struct{}

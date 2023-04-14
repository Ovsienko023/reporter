package repository

import (
	"context"
	"time"
)

const sqlGetVacationPaid = `
	select id,
		   date_from,
		   date_to,
		   status,
		   description,
		   creator_id,
		   created_at,
		   updated_at,
		   payload
    from main.vacations_paid as v
    where id = $1 and exists(select 1
        					 from main.permissions_users_to_objects as pto
            				 where pto.user_id = $2 and
       							   pto.object_id =  v.creator_id)`

func (c *Client) GetVacationPaid(ctx context.Context, msg *GetVacationPaid) (*VacationPaid, error) {
	isAuth, err := c.checkUserPermission(ctx, msg.InvokerId, msg.InvokerId)
	if !isAuth {
		return nil, ErrPermissionDenied
	}

	row, err := c.driver.Query(ctx,
		sqlGetVacationPaid,
		msg.VacationPaidId,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	vacation := &VacationPaid{}

	for row.Next() {
		err := row.Scan(
			&vacation.Id,
			&vacation.DateFrom,
			&vacation.DateTo,
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

type GetVacationPaid struct {
	InvokerId      string `json:"invoker_id,omitempty"`
	VacationPaidId string `json:"vacation_paid_id,omitempty"`
}

type VacationPaid struct {
	Id          *string              `json:"id,omitempty"`
	DateFrom    *time.Time           `json:"date_from,omitempty"`
	DateTo      *time.Time           `json:"date_to,omitempty"`
	Status      *string              `json:"status,omitempty"`
	Description *string              `json:"description,omitempty"`
	CreatorId   *string              `json:"creator_id,omitempty"`
	CreatedAt   *time.Time           `json:"created_at,omitempty"`
	UpdatedAt   *time.Time           `json:"updated_at,omitempty"`
	Payload     *VacationPaidPayload `json:"payload,omitempty"`
}

type VacationPaidPayload struct{}

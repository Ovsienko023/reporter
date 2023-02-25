package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

const sqlDeleteSickLeave = `
	delete from main.sick_leave
    where id = $1 and creator_id = $2`

func (c *Client) DeleteSickLeave(ctx context.Context, msg *DeleteSickLeave) error {
	err := c.isFindSickLeave(ctx, msg.InvokerId, msg.SickLeaveId)
	if errors.Is(err, ErrSickLeaveIdNotFound) {
		return ErrSickLeaveIdNotFound
	}

	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	row, err := transaction.Query(ctx, sqlDeleteSickLeave,
		msg.SickLeaveId,
		msg.InvokerId,
	)
	if err != nil {
		return NewInternalError(err)
	}

	row.Next()
	status := row.CommandTag()
	if status != nil && !status.Delete() {
		return NewInternalError(err)
	}

	if err != nil {
		return NewInternalError(err)
	}

	_ = transaction.Commit(ctx)

	return nil
}

type DeleteSickLeave struct {
	InvokerId   string `json:"invoker_id,omitempty"`
	SickLeaveId string `json:"sick_leave_id,omitempty"`
}

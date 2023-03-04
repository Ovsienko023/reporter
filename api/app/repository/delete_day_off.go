package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

const sqlDeleteDayOff = `
	delete from main.day_off
    where id = $1 and creator_id = $2`

func (c *Client) DeleteDayOff(ctx context.Context, msg *DeleteDayOff) error {
	err := c.isFindDayOff(ctx, msg.InvokerId, msg.DayOffId)
	if errors.Is(err, ErrDayOffIdNotFound) {
		return ErrDayOffIdNotFound
	}

	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	row, err := transaction.Query(ctx, sqlDeleteDayOff,
		msg.DayOffId,
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

type DeleteDayOff struct {
	InvokerId string `json:"invoker_id,omitempty"`
	DayOffId  string `json:"day_off_id,omitempty"`
}

package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

const sqlCreateDayOff = `
    INSERT INTO main.day_off
        (date_from, date_to, status, description, creator_id)
    VALUES
    ($1, $2, $3, $4, $5)
    RETURNING id
`

func (c *Client) CreateDayOff(ctx context.Context, msg *CreateDayOff) (*CreatedDayOff, error) {
	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	row, err := transaction.Query(ctx, sqlCreateDayOff,
		msg.DateFrom,
		msg.DateTo,
		msg.Status,
		msg.Description,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	DayOff := &CreatedDayOff{}

	for row.Next() {
		err = row.Scan(
			&DayOff.Id,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	_ = transaction.Commit(ctx)
	return DayOff, nil
}

type CreateDayOff struct {
	InvokerId   string    `json:"invoker_id,omitempty"`
	DateFrom    time.Time `json:"date_from,omitempty"`
	DateTo      time.Time `json:"date_to,omitempty"`
	Status      string    `json:"status,omitempty"`
	Description string    `json:"description,omitempty"`
}

type CreatedDayOff struct {
	Id string `json:"id,omitempty"`
}

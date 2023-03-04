package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

const sqlCreateSickLeave = `
    INSERT INTO main.sick_leave
        (date_from, date_to, status, description, creator_id)
    VALUES
    ($1, $2, $3, $4, $5)
    RETURNING id
`

func (c *Client) CreateSickLeave(ctx context.Context, msg *CreateSickLeave) (*CreatedSickLeave, error) {
	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	row, err := transaction.Query(ctx, sqlCreateSickLeave,
		msg.DateFrom,
		msg.DateTo,
		msg.Status,
		msg.Description,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	sickLeave := &CreatedSickLeave{}

	for row.Next() {
		err = row.Scan(
			&sickLeave.Id,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	_ = transaction.Commit(ctx)
	return sickLeave, nil
}

type CreateSickLeave struct {
	InvokerId   string    `json:"invoker_id,omitempty"`
	DateFrom    time.Time `json:"date_from,omitempty"`
	DateTo      time.Time `json:"date_to,omitempty"`
	Status      string    `json:"status,omitempty"`
	Description string    `json:"description,omitempty"`
}

type CreatedSickLeave struct {
	Id string `json:"id,omitempty"`
}

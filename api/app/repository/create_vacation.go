package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

const sqlCreateVacation = `
    INSERT INTO main.vacation
        (date, is_paid, status, description, creator_id)
    VALUES
    ($1, $2, $3, $4, $5)
    RETURNING id
`

func (c *Client) CreateVacation(ctx context.Context, msg *CreateVacation) (*CreatedVacation, error) {
	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	row, err := transaction.Query(ctx, sqlCreateVacation,
		msg.Date,
		msg.IsPaid,
		msg.Status,
		msg.Description,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	vacation := &CreatedVacation{}

	for row.Next() {
		err = row.Scan(
			&vacation.Id,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	_ = transaction.Commit(ctx)
	return vacation, nil
}

type CreateVacation struct {
	InvokerId   string    `json:"invoker_id,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	IsPaid      bool      `json:"is_paid,omitempty"`
	Status      string    `json:"status,omitempty"`
	Description string    `json:"description,omitempty"`
}

type CreatedVacation struct {
	Id string `json:"id,omitempty"`
}

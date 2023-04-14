package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

const sqlCreateVacationPaid = `
    INSERT INTO main.vacations_paid
        (date_from, date_to, status, description, creator_id)
    VALUES
    ($1, $2, $3, $4, $5)
    RETURNING id
`

func (c *Client) CreateVacationPaid(ctx context.Context, msg *CreateVacationPaid) (*CreatedVacationPaid, error) {
	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	row, err := transaction.Query(ctx, sqlCreateVacationPaid,
		msg.DateFrom,
		msg.DateTo,
		msg.Status,
		msg.Description,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	vacation := &CreatedVacationPaid{}

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

type CreateVacationPaid struct {
	InvokerId   string    `json:"invoker_id,omitempty"`
	DateFrom    time.Time `json:"date_from,omitempty"`
	DateTo      time.Time `json:"date_to,omitempty"`
	Status      string    `json:"status,omitempty"`
	Description string    `json:"description,omitempty"`
}

type CreatedVacationPaid struct {
	Id string `json:"id,omitempty"`
}

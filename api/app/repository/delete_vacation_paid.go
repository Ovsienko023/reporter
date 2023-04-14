package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

const sqlDeleteVacationPaid = `
	delete from main.vacations_paid
    where id = $1 and creator_id = $2`

func (c *Client) DeleteVacationPaid(ctx context.Context, msg *DeleteVacationPaid) error {
	err := c.isFindVacationPaid(ctx, msg.InvokerId, msg.VacationPaidId)
	if errors.Is(err, ErrVacationIdNotFound) {
		return ErrVacationIdNotFound
	}

	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	row, err := transaction.Query(ctx, sqlDeleteVacationPaid,
		msg.VacationPaidId,
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

type DeleteVacationPaid struct {
	InvokerId      string `json:"invoker_id,omitempty"`
	VacationPaidId string `json:"vacation_paid_id,omitempty"`
}

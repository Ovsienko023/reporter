package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

const sqlDeleteVacationUnpaid = `
	delete from main.vacations_unpaid
    where id = $1 and creator_id = $2`

func (c *Client) DeleteVacationUnpaid(ctx context.Context, msg *DeleteVacationUnpaid) error {
	err := c.isFindVacationUnpaid(ctx, msg.InvokerId, msg.VacationUnpaidId)
	if errors.Is(err, ErrVacationIdNotFound) {
		return ErrVacationIdNotFound
	}

	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	row, err := transaction.Query(ctx, sqlDeleteVacationUnpaid,
		msg.VacationUnpaidId,
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

type DeleteVacationUnpaid struct {
	InvokerId        string `json:"invoker_id,omitempty"`
	VacationUnpaidId string `json:"vacation_unpaid_id,omitempty"`
}

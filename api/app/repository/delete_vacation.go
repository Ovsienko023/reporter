package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

const sqlDeleteVacation = `
	delete from main.vacation
    where id = $1 and creator_id = $2`

func (c *Client) DeleteVacation(ctx context.Context, msg *DeleteVacation) error {
	err := c.isFindVacation(ctx, msg.InvokerId, msg.VacationId)
	if errors.Is(err, ErrVacationIdNotFound) {
		return ErrVacationIdNotFound
	}

	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	row, err := transaction.Query(ctx, sqlDeleteVacation,
		msg.VacationId,
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

type DeleteVacation struct {
	InvokerId  string `json:"invoker_id,omitempty"`
	VacationId string `json:"vacation_id,omitempty"`
}

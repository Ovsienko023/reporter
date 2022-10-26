package database

import (
	"context"
	"github.com/jackc/pgx/v4"
)

const sqlUpdateProfile = `
	update main.users
    set display_name = $2
    where id = $1`

func (c *Client) UpdateProfile(ctx context.Context, msg *UpdateProfile) error {
	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})

	row, err := transaction.Query(ctx, sqlUpdateProfile,
		msg.InvokerId,
		msg.DisplayName,
	)
	if err != nil {
		return NewInternalError(err)
	}

	row.Next()
	status := row.CommandTag()
	if status != nil && !status.Update() {
		return NewInternalError(err)
	}

	if err != nil {
		return NewInternalError(err)
	}

	_ = transaction.Commit(ctx)

	return nil
}

type UpdateProfile struct {
	InvokerId   string `json:"invoker_id,omitempty"`
	DisplayName string `json:"display_name,omitempty"`
}

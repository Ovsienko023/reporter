package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
)

var ErrPermission = errors.New("permission denied")

const sqlAddUserPermission = `
    INSERT INTO main.permissions_users_to_objects
        (user_id, object_type, object_id)
    VALUES
    ($1, $2, $3)
`

func (c *Client) AddObjectToUserPermission(ctx context.Context, msg *AddUserPermission) error {
	if isAdmin, _ := c.checkPermission(ctx, msg.InvokerId); isAdmin != true {
		return ErrPermission
	}

	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
	// todo добавить проверку роли администратора
	rowAdded, err := transaction.Query(ctx, sqlAddUserPermission,
		msg.UserId,
		msg.ObjectType,
		msg.ObjectId,
	)
	if err != nil {
		return NewInternalError(err)
	}

	rowAdded.Next()
	status := rowAdded.CommandTag()

	if status != nil && !status.Insert() {
		return NewInternalError(err)
	}

	_ = transaction.Commit(ctx)
	return nil
}

type AddUserPermission struct {
	InvokerId  string `json:"invoker_id,omitempty"`
	UserId     string `json:"user_id,omitempty"`
	ObjectType string `json:"object_type,omitempty"`
	ObjectId   string `json:"object_id,omitempty"`
}

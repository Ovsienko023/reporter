package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
)

const sqlRemoveObjectFromUserPermission = `
	DELETE 
	FROM main.permissions_users_to_objects as puo
    where puo.user_id = $1 and
          puo.object_id = $2`

func (c *Client) RemoveObjectFromUserPermission(ctx context.Context, msg *RemoveObjectFromUserPermission) error {
	if isAdmin, _ := c.checkPermission(ctx, msg.InvokerId); isAdmin != true {
		return ErrPermission
	}
	// todo Добавить проверки на существование сущьностей
	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("%w: %v", ErrInternal, err)
	}
	// todo добавить проверку роли администратора
	rowRemove, err := transaction.Query(ctx, sqlRemoveObjectFromUserPermission,
		msg.UserId,
		msg.ObjectId,
	)
	if err != nil {
		return NewInternalError(err)
	}

	rowRemove.Next()
	statusRemove := rowRemove.CommandTag()

	if statusRemove != nil && !statusRemove.Delete() {
		return NewInternalError(err)
	}

	_ = transaction.Commit(ctx)
	return nil
}

type RemoveObjectFromUserPermission struct {
	InvokerId string `json:"invoker_id,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	ObjectId  string `json:"object_id,omitempty"`
}

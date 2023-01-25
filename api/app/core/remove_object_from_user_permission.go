package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

// RemoveObjectFromUserPermission возвращает следующие ошибки:
// ErrUnauthorized
// ErrInternal
func (c *Core) RemoveObjectFromUserPermission(ctx context.Context, msg *domain.RemoveObjectFromUserPermissionRequest) error {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return err
	}

	err = c.db.RemoveObjectFromUserPermission(ctx, msg.ToDbRemoveObjectFromUserPermission(invokerId))
	if err != nil {
		return err
	}

	return nil
}

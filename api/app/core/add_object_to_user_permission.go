package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

// AddObjectToUserPermission возвращает следующие ошибки:
// ErrUnauthorized
// ErrInternal
func (c *Core) AddObjectToUserPermission(ctx context.Context, msg *domain.AddObjectToUserPermissionRequest) error {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return err
	}

	err = c.db.AddObjectToUserPermission(ctx, msg.ToDbAddObjectToUserPermission(invokerId))
	if err != nil {
		return err
	}

	return nil
}

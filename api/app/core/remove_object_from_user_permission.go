package core

import (
	"context"
	"errors"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
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
		switch {
		case errors.Is(err, repository.ErrPermissionDenied):
			return ErrPermissionDenied
		default:
			return ErrInternal
		}
	}

	return nil
}

package core

import (
	"context"
	"errors"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
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
		switch {
		case errors.Is(err, repository.ErrPermissionDenied):
			return ErrPermissionDenied
		default:
			return ErrInternal
		}
	}

	return nil
}

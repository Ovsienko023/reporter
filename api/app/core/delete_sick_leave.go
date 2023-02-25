package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
)

func (c *Core) DeleteSickLeave(ctx context.Context, msg *domain.DeleteSickLeaveRequest) error {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return err
	}

	err = c.db.DeleteSickLeave(ctx, msg.ToDbDeleteSickLeave(invokerId))

	if err != nil {
		switch {
		case errors.Is(err, repository.ErrSickLeaveIdNotFound):
			return ErrSickLeaveIdNotFound
		default:
			return fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return nil
}

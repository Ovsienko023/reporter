package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
)

func (c *Core) DeleteDayOff(ctx context.Context, msg *domain.DeleteDayOffRequest) error {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return err
	}

	err = c.db.DeleteDayOff(ctx, msg.ToDbDeleteDayOff(invokerId))

	if err != nil {
		switch {
		case errors.Is(err, repository.ErrDayOffIdNotFound):
			return ErrDayOffIdNotFound
		default:
			return fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return nil
}

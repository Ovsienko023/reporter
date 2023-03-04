package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
)

func (c *Core) GetDayOff(ctx context.Context, msg *domain.GetDayOffRequest) (*domain.GetDayOffResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.GetDayOff(ctx, msg.ToDbGetDayOff(invokerId))

	if err != nil {
		switch {
		case errors.Is(err, repository.ErrDayOffIdNotFound):
			return nil, ErrDayOffIdNotFound
		case errors.Is(err, repository.ErrPermissionDenied):
			return nil, ErrPermissionDenied
		default:
			return nil, fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return domain.FromGetDayOffResponse(result), nil
}

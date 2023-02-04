package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
)

func (c *Core) GetSickLeave(ctx context.Context, msg *domain.GetSickLeaveRequest) (*domain.GetSickLeaveResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.GetSickLeave(ctx, msg.ToDbGetSickLeave(invokerId))

	if err != nil {
		switch {
		case errors.Is(err, repository.ErrSickLeaveIdNotFound):
			return nil, ErrSickLeaveIdNotFound
		case errors.Is(err, repository.ErrPermissionDenied):
			return nil, ErrPermissionDenied
		default:
			return nil, fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return domain.FromGetSickLeaveResponse(result), nil
}

package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
)

func (c *Core) GetVacationUnpaid(ctx context.Context, msg *domain.GetVacationUnpaidRequest) (*domain.GetVacationUnpaidResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.GetVacationUnpaid(ctx, msg.ToDbGetVacationUnpaid(invokerId))

	if err != nil {
		switch {
		case errors.Is(err, repository.ErrVacationIdNotFound):
			return nil, ErrVacationIdNotFound
		case errors.Is(err, repository.ErrPermissionDenied):
			return nil, ErrPermissionDenied
		default:
			return nil, fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return domain.FromGetVacationUnpaidResponse(result), nil
}

package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
)

func (c *Core) GetVacation(ctx context.Context, msg *domain.GetVacationRequest) (*domain.GetVacationResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.GetVacation(ctx, msg.ToDbGetVacation(invokerId))

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

	return domain.FromGetVacationResponse(result), nil
}

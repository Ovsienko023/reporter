package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

// CreateVacationUnpaid возвращает следующие ошибки:
// ErrUnauthorized
// ErrInternal
func (c *Core) CreateVacationUnpaid(ctx context.Context, msg *domain.CreateVacationUnpaidRequest) (*domain.CreateVacationUnpaidResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.CreateVacationUnpaid(ctx, msg.ToDbCreateVacationUnpaid(invokerId))
	if err != nil {
		return nil, err
	}

	return domain.FromCreateVacationUnpaidResponse(result), nil
}

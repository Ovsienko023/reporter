package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

// CreateVacationPaid возвращает следующие ошибки:
// ErrUnauthorized
// ErrInternal
func (c *Core) CreateVacationPaid(ctx context.Context, msg *domain.CreateVacationPaidRequest) (*domain.CreateVacationPaidResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.CreateVacationPaid(ctx, msg.ToDbCreateVacationPaid(invokerId))
	if err != nil {
		return nil, err
	}

	return domain.FromCreateVacationPaidResponse(result), nil
}

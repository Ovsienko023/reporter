package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

// CreateVacation возвращает следующие ошибки:
// ErrUnauthorized
// ErrInternal
func (c *Core) CreateVacation(ctx context.Context, msg *domain.CreateVacationRequest) (*domain.CreateVacationResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.CreateVacation(ctx, msg.ToDbCreateVacation(invokerId))
	if err != nil {
		return nil, err
	}

	return domain.FromCreateVacationResponse(result), nil
}

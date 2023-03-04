package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

// CreateDayOff возвращает следующие ошибки:
// ErrUnauthorized
// ErrInternal
func (c *Core) CreateDayOff(ctx context.Context, msg *domain.CreateDayOffRequest) (*domain.CreateDayOffResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.CreateDayOff(ctx, msg.ToDbCreateDayOff(invokerId))
	if err != nil {
		return nil, err
	}

	return domain.FromCreateDayOffResponse(result), nil
}

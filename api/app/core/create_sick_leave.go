package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

// CreateSickLeave возвращает следующие ошибки:
// ErrUnauthorized
// ErrInternal
func (c *Core) CreateSickLeave(ctx context.Context, msg *domain.CreateSickLeaveRequest) (*domain.CreateSickLeaveResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.CreateSickLeave(ctx, msg.ToDbCreateSickLeave(invokerId))
	if err != nil {
		return nil, err
	}

	return domain.FromCreateSickLeaveResponse(result), nil
}

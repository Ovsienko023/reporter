package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

// CreateReport возвращает следующие ошибки:
// ErrUnauthorized
// ErrInternal
func (c *Core) CreateReport(ctx context.Context, msg *domain.CreateReportRequest) (*domain.CreateReportResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.CreateReport(ctx, msg.ToDbCreateReport(invokerId))
	if err != nil {
		return nil, err
	}

	return domain.FromCreateReportResponse(result), nil
}

package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
)

func (c *Core) CreateReport(ctx context.Context, msg *domain.CreateReportRequest) (*domain.CreateReportResponse, error) {
	result, err := c.db.CreateReport(ctx, msg.ToDBCreateReport())
	if err != nil {
		return nil, err
	}

	resp := domain.CreateReportResponse{
		Id: result.Id,
	}

	return &resp, nil
}

package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

func (c *Core) GetReport(ctx context.Context, msg *domain.GetReportRequest) (*domain.GetReportResponse, error) {
	result, err := c.db.GetReport(ctx, msg.ToDbGetReport())

	if err != nil {
		switch {
		case errors.Is(err, database.ErrReportIdNotFound):
			return nil, ErrReportIdNotFound
		default:
			return nil, fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return domain.FromGetReportResponse(result), nil
}

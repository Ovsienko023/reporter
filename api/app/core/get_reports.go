package core

import (
	"context"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
)

func (c *Core) GetReports(ctx context.Context, msg *domain.GetReportsRequest) (*domain.GetReportsResponse, error) {
	result, cnt, err := c.db.GetReports(ctx, msg.ToDbGetReports())
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	return domain.FromGetReportsResponse(result, cnt), nil
}

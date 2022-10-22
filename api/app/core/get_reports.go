package core

import (
	"context"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
)

func (c *Core) GetReports(ctx context.Context, msg *domain.GetReportsRequest) (*domain.GetReportsResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, cnt, err := c.db.GetReports(ctx, msg.ToDbGetReports(invokerId))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	return domain.FromGetReportsResponse(result, cnt), nil
}

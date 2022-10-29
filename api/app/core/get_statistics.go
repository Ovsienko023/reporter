package core

import (
	"context"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
)

func (c *Core) GetStatistics(ctx context.Context, msg *domain.GetStatisticsRequest) (*domain.GetStatisticsResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, err := c.db.GetStatistics(ctx, msg.ToDbGetStatistics(invokerId))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	return domain.FromGetStatisticsResponse(result), nil
}

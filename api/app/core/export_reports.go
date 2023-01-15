package core

import (
	"context"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
)

func (c *Core) ExportReports(ctx context.Context, msg *domain.ExportReportsRequest) (*domain.ExportReportsResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, _, err := c.db.GetReports(ctx, msg.ToDbGetReports(invokerId))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	return new(domain.ExportReportsResponse).From(result), nil
}

package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

func (c *Core) DeleteReport(ctx context.Context, msg *domain.DeleteReportRequest) error {
	err := c.db.DeleteReport(ctx, msg.ToDbDeleteReport())

	if err != nil {
		switch {
		case errors.Is(err, database.ErrReportIdNotFound):
			return ErrReportIdNotFound
		default:
			return fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return nil
}

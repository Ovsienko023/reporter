package core

import (
	"context"
	"errors"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

func (c *Core) UpdateReport(ctx context.Context, msg *domain.UpdateReportRequest) error {
	err := c.db.UpdateReport(ctx, msg.ToDbUpdateReport())

	if err != nil {
		switch {
		case errors.Is(err, database.ErrReportIdNotFound):
			return ErrReportIdNotFound
		}
		return err
	}

	return nil
}

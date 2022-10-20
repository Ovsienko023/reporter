package core

import (
	"context"
	"errors"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

func (c *Core) DeleteReport(ctx context.Context, msg *domain.DeleteReportRequest) error {
	message := database.DeleteReport{
		InvokerId: msg.InvokerId,
		ReportId:  msg.ReportId,
	}

	err := c.db.DeleteReport(ctx, &message)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrReportIdNotFound):
			return ErrReportIdNotFound
		}
		return err
	}

	return nil
}

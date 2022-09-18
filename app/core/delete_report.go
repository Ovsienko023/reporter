package core

import (
	"context"
	"errors"
	"github.com/Ovsienko023/reporter/app/domain"
	database2 "github.com/Ovsienko023/reporter/infrastructure/database"
)

func (c *Core) DeleteReport(ctx context.Context, msg *domain.DeleteReportRequest) error {
	systemUser := c.repo.GetSystemUser()

	message := database2.DeleteReport{
		InvokerId: *systemUser.UserId,
		ReportId:  msg.ReportId,
	}

	err := c.repo.DeleteReport(ctx, &message)
	if err != nil {
		switch {
		case errors.Is(err, database2.ErrReportIdNotFound):
			return ErrReportIdNotFound
		}
		return err
	}

	return nil
}

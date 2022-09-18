package core

import (
	"context"
	"errors"
	"github.com/Ovsienko023/reporter/app/domain"
	database2 "github.com/Ovsienko023/reporter/infrastructure/database"
)

func (c *Core) UpdateReport(ctx context.Context, msg *domain.UpdateReportRequest) error {
	systemUser := c.repo.GetSystemUser()

	message := database2.UpdateReport{
		InvokerId: *systemUser.UserId,
		ReportId:  msg.ReportId,
		Title:     msg.Title,
		StartTime: msg.StartTime,
		EndTime:   msg.EndTime,
		BreakTime: msg.BreakTime,
		WorkTime:  msg.WorkTime,
		Body:      msg.Body,
	}

	err := c.repo.UpdateReport(ctx, &message)
	if err != nil {
		switch {
		case errors.Is(err, database2.ErrReportIdNotFound):
			return ErrReportIdNotFound
		}
		return err
	}

	return nil
}

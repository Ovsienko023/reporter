package core

import (
	"context"
	"errors"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
	"time"
)

func (c *Core) UpdateReport(ctx context.Context, msg *domain.UpdateReportRequest) error {
	message := database.UpdateReport{
		InvokerId: msg.InvokerId,
		ReportId:  msg.ReportId,
		Title:     msg.Title,
		Date:      time.Unix(msg.Date, 0),
		StartTime: time.Unix(msg.StartTime, 0),
		EndTime:   time.Unix(msg.EndTime, 0),
		BreakTime: time.Unix(msg.BreakTime, 0),
		WorkTime:  time.Unix(msg.WorkTime, 0),
		Body:      msg.Body,
	}

	err := c.db.UpdateReport(ctx, &message)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrReportIdNotFound):
			return ErrReportIdNotFound
		}
		return err
	}

	return nil
}

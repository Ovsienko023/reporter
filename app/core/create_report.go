package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
	"time"
)

func (c *Core) CreateReport(ctx context.Context, msg *domain.CreateReportRequest) (*domain.CreateReportResponse, error) {
	systemUser := c.db.GetSystemUser()

	message := database.CreateReport{
		InvokerId: *systemUser.UserId,
		Title:     msg.Title,
		Date:      time.Unix(msg.Date, 0).UTC(),
		StartTime: time.Unix(msg.StartTime, 0).UTC(),
		EndTime:   time.Unix(msg.EndTime, 0).UTC(),
		BreakTime: time.Unix(msg.BreakTime, 0).UTC(),
		WorkTime:  time.Unix(msg.WorkTime, 0).UTC(),
		Body:      msg.Body,
	}
	result, err := c.db.CreateReport(ctx, &message)
	if err != nil {
		return nil, err
	}

	resp := domain.CreateReportResponse{
		Id: result.Id,
	}

	return &resp, nil
}

package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

func (c *Core) CreateReport(ctx context.Context, msg *domain.CreateReportRequest) (*domain.CreateReportResponse, error) {
	systemUser := c.repo.GetSystemUser()

	message := database.CreateReport{
		InvokerId: *systemUser.UserId,
		Title:     msg.Title,
		Date:      msg.Date,
		StartTime: msg.StartTime,
		EndTime:   msg.EndTime,
		BreakTime: msg.BreakTime,
		WorkTime:  msg.WorkTime,
		Body:      msg.Body,
	}
	result, err := c.repo.CreateReport(ctx, &message)
	if err != nil {
		return nil, err
	}

	resp := domain.CreateReportResponse{
		Id: result.Id,
	}

	return &resp, nil
}

package core

import (
	"context"
	"errors"
	"github.com/Ovsienko023/reporter/app/domain"
	database2 "github.com/Ovsienko023/reporter/infrastructure/database"
)

func (c *Core) GetReport(ctx context.Context, msg *domain.GetReportRequest) (*domain.GetReportResponse, error) {
	systemUser := c.repo.GetSystemUser()

	message := database2.GetReport{
		InvokerId: *systemUser.UserId,
		ReportId:  msg.ReportId,
	} // domain.GetReportRequest to database.GetReport

	result, err := c.repo.GetReport(ctx, &message)
	if err != nil {
		switch {
		case errors.Is(err, database2.ErrReportIdNotFound):
			return nil, ErrReportIdNotFound
		}
		return nil, err
	}

	resp := domain.GetReportResponse{
		Report: &domain.Report{
			Id:        result.Id,
			Title:     result.Title,
			CreatorId: result.CreatorId,
			CreatedAt: result.CreatedAt,
			UpdatedAt: result.UpdatedAt,
			DeletedAt: result.DeletedAt,
			StartTime: result.StartTime,
			EndTime:   result.EndTime,
			BreakTime: result.BreakTime,
			WorkTime:  result.WorkTime,
			Body:      result.Body,
		},
	}

	return &resp, nil
}

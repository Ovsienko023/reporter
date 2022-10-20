package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

func (c *Core) GetReport(ctx context.Context, msg *domain.GetReportRequest) (*domain.GetReportResponse, error) {
	message := database.GetReport{
		InvokerId: msg.InvokerId,
		ReportId:  msg.ReportId,
	}

	result, err := c.db.GetReport(ctx, &message)
	if err != nil {
		switch {
		case errors.Is(err, database.ErrReportIdNotFound):
			return nil, ErrReportIdNotFound
		}
		fmt.Println("LOG: ", err) // todo add logger
		return nil, ErrInternal
	}

	resp := domain.GetReportResponse{
		Report: &domain.Report{
			Id:        result.Id,
			Title:     result.Title,
			Date:      ptr.Int64(result.Date.Unix()),
			CreatorId: result.CreatorId,
			CreatedAt: ptr.Int64(result.CreatedAt.Unix()),
			UpdatedAt: ptr.Int64(result.UpdatedAt.Unix()),
			StartTime: ptr.Int64(result.StartTime.Unix()),
			EndTime:   ptr.Int64(result.EndTime.Unix()),
			BreakTime: ptr.Int64(result.BreakTime.Unix()),
			WorkTime:  ptr.Int64(result.WorkTime.Unix()),
			Body:      result.Body,
		},
	}

	return &resp, nil
}

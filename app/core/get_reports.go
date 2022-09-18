package core

import (
	"context"
	domain2 "github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

func (c *Core) GetReports(ctx context.Context, msg *domain2.GetReportsRequest) (*domain2.GetReportsResponse, error) {
	systemUser := c.repo.GetSystemUser()

	message := database.GetReports{
		InvokerId: *systemUser.UserId,
	}

	result, cnt, err := c.repo.GetReports(ctx, &message)
	if err != nil {
		return nil, err
	}

	reports := make([]domain2.Report, 0, len(result.Reports))

	for _, obj := range result.Reports {
		item := domain2.Report{
			Id:        obj.Id,
			Title:     obj.Title,
			CreatorId: obj.CreatorId,
			CreatedAt: obj.CreatedAt,
			UpdatedAt: obj.UpdatedAt,
			DeletedAt: obj.DeletedAt,
			StartTime: obj.StartTime,
			EndTime:   obj.EndTime,
			BreakTime: obj.BreakTime,
			WorkTime:  obj.WorkTime,
			Body:      obj.Body,
		}
		reports = append(reports, item)
	}

	return &domain2.GetReportsResponse{
		Count:   cnt,
		Reports: reports,
	}, nil
}

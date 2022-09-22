package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
)

func (c *Core) GetReports(ctx context.Context, msg *domain.GetReportsRequest) (*domain.GetReportsResponse, error) {
	systemUser := c.repo.GetSystemUser()

	message := database.GetReports{
		InvokerId: *systemUser.UserId,
	}

	result, cnt, err := c.repo.GetReports(ctx, &message)
	if err != nil {
		return nil, err
	}

	reports := make([]domain.Report, 0, len(result.Reports))

	for _, obj := range result.Reports {
		item := domain.Report{
			Id:        obj.Id,
			Title:     obj.Title,
			Date:      obj.Date,
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

	return &domain.GetReportsResponse{
		Count:   cnt,
		Reports: reports,
	}, nil
}

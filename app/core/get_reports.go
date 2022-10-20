package core

import (
	"context"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/infrastructure/database"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
)

func (c *Core) GetReports(ctx context.Context, msg *domain.GetReportsRequest) (*domain.GetReportsResponse, error) {
	message := database.GetReports{
		InvokerId: msg.InvokerId,
	}

	result, cnt, err := c.db.GetReports(ctx, &message)
	if err != nil {
		return nil, err
	}

	reports := make([]domain.ReportItem, 0, len(result))

	for _, obj := range result {
		var deletedAt *int64
		if obj.DeletedAt != nil {
			deletedAt = ptr.Int64(obj.DeletedAt.Unix())
		}
		item := domain.ReportItem{
			Id:        obj.Id,
			Title:     obj.Title,
			Date:      ptr.Int64(obj.Date.Unix()),
			CreatorId: obj.CreatorId,
			CreatedAt: ptr.Int64(obj.CreatedAt.Unix()),
			UpdatedAt: ptr.Int64(obj.UpdatedAt.Unix()),
			DeletedAt: deletedAt,
			StartTime: ptr.Int64(obj.StartTime.Unix()),
			EndTime:   ptr.Int64(obj.EndTime.Unix()),
			BreakTime: ptr.Int64(obj.BreakTime.Unix()),
			WorkTime:  ptr.Int64(obj.WorkTime.Unix()),
			Body:      obj.Body,
		}
		reports = append(reports, item)
	}

	return &domain.GetReportsResponse{
		Count:   cnt,
		Reports: reports,
	}, nil
}

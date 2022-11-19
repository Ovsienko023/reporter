package core

import (
	"context"
	"errors"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
)

func (c *Core) UpdateReport(ctx context.Context, msg *domain.UpdateReportRequest) error {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return err
	}

	err = c.db.UpdateReport(ctx, msg.ToDbUpdateReport(invokerId))
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrReportIdNotFound):
			return ErrReportIdNotFound
		}
		return err
	}

	return nil
}

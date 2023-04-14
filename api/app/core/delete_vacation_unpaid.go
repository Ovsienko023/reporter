package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
)

func (c *Core) DeleteVacationUnpaid(ctx context.Context, msg *domain.DeleteVacationUnpaidRequest) error {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return err
	}

	err = c.db.DeleteVacationUnpaid(ctx, msg.ToDbDeleteVacationUnpaid(invokerId))

	if err != nil {
		switch {
		case errors.Is(err, repository.ErrVacationIdNotFound):
			return ErrVacationIdNotFound
		default:
			return fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return nil
}

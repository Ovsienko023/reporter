package core

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/repository"
)

func (c *Core) GetCalendarEvents(ctx context.Context, msg *domain.GetCalendarEventsRequest) (*domain.GetCalendarEventsResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, cnt, err := c.db.GetCalendarEvents(ctx, msg.ToDbGetCalendarEvents(invokerId))
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrUserIdNotFound):
			return nil, ErrUserIdFromAllowedToNotFound
		case errors.Is(err, repository.ErrUnauthorized):
			return nil, ErrUnauthorized
		default:
			return nil, fmt.Errorf("%w: %v", ErrInternal, err)
		}
	}

	return new(domain.GetCalendarEventsResponse).FromGetCalendarEvents(result, cnt), nil
}

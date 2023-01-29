package core

import (
	"context"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain"
)

func (c *Core) GetCalendarEvents(ctx context.Context, msg *domain.GetCalendarEventsRequest) (*domain.GetCalendarEventsResponse, error) {
	invokerId, err := c.authorize(msg.Token)
	if err != nil {
		return nil, err
	}

	result, cnt, err := c.db.GetCalendarEvents(ctx, msg.ToDbGetCalendarEvents(invokerId))
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	return new(domain.GetCalendarEventsResponse).FromGetCalendarEvents(result, cnt), nil
}

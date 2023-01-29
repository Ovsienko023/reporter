package repository

import (
	"context"
	"time"
)

const (
	sqlGetCalendarEvents = `
	with tab as (
		select id,
			   event_type,
			   date
		from main.events
		where events.user_id = $1 and 
			($2::timestamp is null and $3::timestamp is null or 
				date >= $2::timestamp and 
				date <= $3::timestamp)
	)
	select (select count(*) from tab) as count,
			r.id                      as event_id,
       		r.event_type              as event_type,
       		r.date                    as date
	from tab as r
	limit $4 offset $4 * ($5 - 1)`
)

func (c *Client) GetCalendarEvents(ctx context.Context, msg *GetCalendarEvents) ([]CalendarEvent, *int64, error) {
	row, err := c.driver.Query(ctx, sqlGetCalendarEvents,
		msg.InvokerId,
		msg.DateFrom,
		msg.DateTo,
		msg.PageSize,
		msg.Page,
	)
	if err != nil {
		return nil, nil, NewInternalError(err)
	}

	var count *int64
	events := make([]CalendarEvent, 0, 0)

	for row.Next() {
		event := CalendarEvent{}
		err := row.Scan(
			&count,
			&event.Id,
			&event.EventType,
			&event.Date,
		)
		if err != nil {
			return nil, nil, NewInternalError(err)
		}
		events = append(events, event)
	}

	return events, count, nil
}

type GetCalendarEvents struct {
	InvokerId string     `json:"invoker_id,omitempty"`
	DateFrom  *time.Time `json:"date_from,omitempty"`
	DateTo    *time.Time `json:"date_to,omitempty"`
	Page      *int       `json:"page,omitempty"`
	PageSize  *int       `json:"page_size,omitempty"`
}

type CalendarEvent struct {
	Id        *string    `json:"id,omitempty"`
	EventType *string    `json:"event_type,omitempty"`
	Date      *time.Time `json:"date,omitempty"`
}

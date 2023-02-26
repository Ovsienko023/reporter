package repository

import (
	"context"
	"time"
)

const sqlGetCalendarEvents = `
	select error,
	       count,
	       event_id,
	       event_type,
	       date
	from main.get_calendar(
	    _invoker_id := $1, 
	    _date_from := $2, 
	    _date_to := $3, 
	    _page := $4, 
	    _page_size := $5, 
	    _allowed_to := $6
	)`

func (c *Client) GetCalendarEvents(ctx context.Context, msg *GetCalendarEvents) ([]CalendarEvent, *int64, error) {
	row, err := c.driver.Query(ctx, sqlGetCalendarEvents,
		msg.InvokerId,
		msg.DateFrom,
		msg.DateTo,
		msg.Page,
		msg.PageSize,
		msg.AllowedTo,
	)
	if err != nil {
		return nil, nil, NewInternalError(err)
	}

	var (
		count    *int64
		queryErr []byte
	)

	events := make([]CalendarEvent, 0, 0)

	for row.Next() {
		event := CalendarEvent{}
		err := row.Scan(
			&queryErr,
			&count,
			&event.Id,
			&event.EventType,
			&event.Date,
		)
		if err != nil {
			return nil, nil, NewInternalError(err)
		}
		if queryErr != nil {
			if err = AnalyzeError(queryErr); err != nil {
				return nil, count, err
			}
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
	AllowedTo *string    `json:"allowed_to,omitempty"`
}

type CalendarEvent struct {
	Id        *string    `json:"id,omitempty"`
	EventType *string    `json:"event_type,omitempty"`
	Date      *time.Time `json:"date,omitempty"`
}

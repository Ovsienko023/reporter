package domain

import (
	"github.com/Ovsienko023/reporter/app/repository"
	"github.com/Ovsienko023/reporter/infrastructure/utils/ptr"
	"time"
)

type GetCalendarEventsRequest struct {
	Token     string     `json:"token,omitempty" swaggerignore:"true"`
	DateFrom  *time.Time `json:"date_from,omitempty"`
	DateTo    *time.Time `json:"date_to,omitempty"`
	Page      *int       `json:"page,omitempty"`
	PageSize  *int       `json:"page_size,omitempty"`
	AllowedTo *string    `json:"allowed_to,omitempty"`
}

func (r *GetCalendarEventsRequest) ToDbGetCalendarEvents(invokerId string) *repository.GetCalendarEvents {
	return &repository.GetCalendarEvents{
		InvokerId: invokerId,
		DateFrom:  r.DateFrom,
		DateTo:    r.DateTo,
		Page:      r.Page,
		PageSize:  r.PageSize,
		AllowedTo: r.AllowedTo,
	}
}

type GetCalendarEventsResponse struct {
	Count  int64           `json:"count,omitempty"`
	Events []CalendarEvent `json:"events,omitempty"`
}

type CalendarEvent struct {
	Id        *string `json:"id,omitempty"`
	EventType *string `json:"event_type,omitempty"`
	DateFrom  *int64  `json:"date_from,omitempty"`
	DateTo    *int64  `json:"date_to,omitempty"`
}

func (e *GetCalendarEventsResponse) FromGetCalendarEvents(events []repository.CalendarEvent, cnt *int64) *GetCalendarEventsResponse {
	calendarEvents := make([]CalendarEvent, 0, len(events))

	for _, event := range events {
		item := CalendarEvent{
			Id:        event.Id,
			EventType: event.EventType,
			DateFrom:  ptr.Int64(event.DateFrom.Unix()),
			DateTo:    ptr.Int64(event.DateTo.Unix()),
		}
		calendarEvents = append(calendarEvents, item)
	}

	result := &GetCalendarEventsResponse{
		Events: calendarEvents,
	}

	if cnt != nil {
		result.Count = *cnt
	}

	return result
}

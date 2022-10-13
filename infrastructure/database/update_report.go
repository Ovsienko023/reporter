package database

import (
	"context"
	"errors"
	"time"
)

const sqlUpdateReport = `
	update main.reports
    set title = $2, 
        date = $3, 
        start_time = $4, 
        end_time = $5, 
        break_time = $6, 
        work_time = $7, 
        body = $8, 
        updated_at = now()
    where id = $1
    returning 1;`

func (c *Client) UpdateReport(ctx context.Context, msg *UpdateReport) error {
	_, err := c.GetReport(ctx, &GetReport{ReportId: msg.ReportId})
	if errors.Is(err, ErrReportIdNotFound) {
		return ErrReportIdNotFound
	}

	row, err := c.driver.Query(ctx, sqlUpdateReport,
		msg.ReportId,
		msg.Title,
		msg.Date,
		msg.StartTime,
		msg.EndTime,
		msg.BreakTime,
		msg.WorkTime,
		msg.Body,
	)
	if err != nil {
		return NewInternalError(err)
	}

	var updated *int

	row.Next()
	err = row.Scan(&updated)
	if err != nil {
		return NewInternalError(err)
	}

	if updated == nil {
		return ErrReportUpdated
	}

	return nil
}

type UpdateReport struct {
	InvokerId string    `json:"invoker_id,omitempty"`
	ReportId  string    `json:"id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Date      time.Time `json:"date,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
	BreakTime time.Time `json:"break_time,omitempty"`
	WorkTime  time.Time `json:"work_time,omitempty"`
	Body      string    `json:"body,omitempty"`
}

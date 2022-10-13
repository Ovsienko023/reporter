package database

import (
	"context"
	"time"
)

const sqlCreateReport = `
    INSERT INTO main.reports
        (title, date, start_time, end_time, break_time, work_time, body, creator_id)
    VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id
`

func (c *Client) CreateReport(ctx context.Context, msg *CreateReport) (*CreatedReport, error) {
	row, err := c.driver.Query(ctx, sqlCreateReport,
		msg.Title,
		msg.Date,
		msg.StartTime,
		msg.EndTime,
		msg.BreakTime,
		msg.WorkTime,
		msg.Body,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	report := &CreatedReport{}

	for row.Next() {
		err = row.Scan(
			&report.Id,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	return report, nil
}

type CreateReport struct {
	InvokerId string    `json:"invoker_id,omitempty"`
	Date      time.Time `json:"date,omitempty"`
	Title     string    `json:"title,omitempty"`
	StartTime time.Time `json:"start_time,omitempty"`
	EndTime   time.Time `json:"end_time,omitempty"`
	BreakTime time.Time `json:"break_time,omitempty"`
	WorkTime  time.Time `json:"work_time,omitempty"`
	Body      string    `json:"body,omitempty"`
}

type CreatedReport struct {
	Id string `json:"id,omitempty"`
}

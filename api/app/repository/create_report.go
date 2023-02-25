package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

const sqlCreateReport = `
    INSERT INTO main.reports
        (display_name, date, start_time, end_time, break_time, work_time, body, creator_id)
    VALUES
    ($1, $2, $3, $4, $5, $6, $7, $8)
    RETURNING id
`

func (c *Client) CreateReport(ctx context.Context, msg *CreateReport) (*CreatedReport, error) {
	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrInternal, err)
	}

	row, err := transaction.Query(ctx, sqlCreateReport,
		msg.DisplayName,
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

	_ = transaction.Commit(ctx)
	return report, nil
}

type CreateReport struct {
	InvokerId   string    `json:"invoker_id,omitempty"`
	Date        time.Time `json:"date,omitempty"`
	DisplayName string    `json:"display_name,omitempty"`
	StartTime   int64     `json:"start_time,omitempty"`
	EndTime     int64     `json:"end_time,omitempty"`
	BreakTime   int64     `json:"break_time,omitempty"`
	WorkTime    int64     `json:"work_time,omitempty"`
	Body        string    `json:"body,omitempty"`
}

type CreatedReport struct {
	Id string `json:"id,omitempty"`
}

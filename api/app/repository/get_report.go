package repository

import (
	"context"
	"time"
)

const sqlGetReport = `
	select id,
	       display_name,
	       state,
		   date,
		   start_time,
		   end_time,
		   break_time,
		   work_time,
		   body,
		   creator_id,
		   created_at,
		   updated_at
    from main.reports
    where id = $1 and creator_id = $2`

func (c *Client) GetReport(ctx context.Context, msg *GetReport) (*Report, error) {
	row, err := c.driver.Query(ctx, sqlGetReport, msg.ReportId, msg.InvokerId)
	if err != nil {
		return nil, NewInternalError(err)
	}

	report := &Report{}

	for row.Next() {
		err := row.Scan(
			&report.Id,
			&report.DisplayName,
			&report.State,
			&report.Date,
			&report.StartTime,
			&report.EndTime,
			&report.BreakTime,
			&report.WorkTime,
			&report.Body,
			&report.CreatorId,
			&report.CreatedAt,
			&report.UpdatedAt,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
	}

	if report.Id == nil {
		return nil, ErrReportIdNotFound
	}

	return report, nil
}

type GetReport struct {
	InvokerId string `json:"invoker_id,omitempty"`
	ReportId  string `json:"report_id,omitempty"`
}

type Report struct {
	Id          *string    `json:"id,omitempty"`
	DisplayName *string    `json:"display_name,omitempty"`
	State       *string    `json:"state,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	StartTime   *int64     `json:"start_time,omitempty"`
	EndTime     *int64     `json:"end_time,omitempty"`
	BreakTime   *int64     `json:"break_time,omitempty"`
	WorkTime    *int64     `json:"work_time,omitempty"`
	Body        *string    `json:"body,omitempty"`
	CreatorId   *string    `json:"creator_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

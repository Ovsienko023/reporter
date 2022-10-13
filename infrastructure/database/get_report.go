package database

import (
	"context"
	"time"
)

const sqlGetReport = `
	select id,
	       title,
		   date,
		   start_time,
		   end_time,
		   break_time,
		   work_time,
		   body,
		   creator_id,
		   created_at,
		   updated_at,
		   deleted_at 
    from main.reports
    where id = $1` // todo del deleted_at

func (c *Client) GetReport(ctx context.Context, msg *GetReport) (*Report, error) {
	row, err := c.driver.Query(ctx, sqlGetReport, msg.ReportId)
	if err != nil {
		return nil, NewInternalError(err)
	}

	report := &Report{}

	for row.Next() {
		err := row.Scan(
			&report.Id,
			&report.Title,
			&report.Date,
			&report.StartTime,
			&report.EndTime,
			&report.BreakTime,
			&report.WorkTime,
			&report.Body,
			&report.CreatorId,
			&report.CreatedAt,
			&report.UpdatedAt,
			&report.DeletedAt,
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
	Id        *string    `json:"id,omitempty"`
	Title     *string    `json:"title,omitempty"`
	Date      *time.Time `json:"date,omitempty"`
	StartTime *time.Time `json:"start_time,omitempty"`
	EndTime   *time.Time `json:"end_time,omitempty"`
	BreakTime *time.Time `json:"break_time,omitempty"`
	WorkTime  *time.Time `json:"work_time,omitempty"`
	Body      *string    `json:"body,omitempty"`
	CreatorId *string    `json:"creator_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

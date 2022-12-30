package repository

import (
	"context"
	"time"
)

const (
	sqlGetReports = `
	with tab as (
		select id,
			   display_name,
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
		inner join main.reports_to_users rtu on reports.id = rtu.report_id
		where rtu.user_id = $1 and 
			($2::timestamp is null and $3::timestamp is null or 
				date >= $2::timestamp and 
				date <= $3::timestamp)
	)
	select (select count(*) from tab) as count,
			r.id                      as report_id,
       		r.display_name            as display_name,
       		r.date                    as date,
       		r.start_time              as start_time,
			r.end_time                as end_time,
			r.break_time              as break_time,
			r.work_time               as work_time,
			r.body                    as body,
			r.creator_id              as creator_id,
			r.created_at              as created_at,
			r.updated_at              as updated_at,
			r.deleted_at              as deleted_at
	from tab as r
	limit $4 offset $4 * ($5 - 1)`
)

func (c *Client) GetReports(ctx context.Context, msg *GetReports) ([]ReportItem, *int, error) {
	row, err := c.driver.Query(ctx, sqlGetReports,
		msg.InvokerId,
		msg.DateFrom,
		msg.DateTo,
		msg.PageSize,
		msg.Page,
	)
	if err != nil {
		return nil, nil, NewInternalError(err)
	}

	var count int
	reports := make([]ReportItem, 0, 0)

	for row.Next() {
		report := ReportItem{}
		err := row.Scan(
			&count,
			&report.Id,
			&report.DisplayName,
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
			return nil, nil, NewInternalError(err)
		}
		reports = append(reports, report)
	}

	return reports, &count, nil
}

type GetReports struct {
	InvokerId string     `json:"invoker_id,omitempty"`
	DateFrom  *time.Time `json:"date_from,omitempty"`
	DateTo    *time.Time `json:"date_to,omitempty"`
	Page      *int       `json:"page,omitempty"`
	PageSize  *int       `json:"page_size,omitempty"`
}

type ReportItem struct {
	Id          *string    `json:"id,omitempty"`
	DisplayName *string    `json:"display_name,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	StartTime   *int64     `json:"start_time,omitempty"`
	EndTime     *int64     `json:"end_time,omitempty"`
	BreakTime   *int64     `json:"break_time,omitempty"`
	WorkTime    *int64     `json:"work_time,omitempty"`
	Body        *string    `json:"body,omitempty"`
	CreatorId   *string    `json:"creator_id,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

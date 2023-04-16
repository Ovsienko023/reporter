package repository

import (
	"context"
	"time"
)

const (
	sqlGetReports = `
	select error,
           count,
           id,
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
	from main.get_reports(
	    _invoker_id := $1, 
	    _date_from := $2, 
	    _date_to := $3, 
	    _page := $4, 
	    _page_size := $5,
	    _allowed_to := $6
	)`
)

func (c *Client) GetReports(ctx context.Context, msg *GetReports) ([]ReportItem, *int, error) {
	row, err := c.driver.Query(ctx, sqlGetReports,
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
		count    *int
		queryErr []byte
	)

	reports := make([]ReportItem, 0, 0)

	for row.Next() {
		report := ReportItem{}
		err := row.Scan(
			&queryErr,
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
		if queryErr != nil {
			if err = AnalyzeError(queryErr); err != nil {
				return nil, count, err
			}
		}
		reports = append(reports, report)
	}

	return reports, count, nil
}

type GetReports struct {
	InvokerId string     `json:"invoker_id,omitempty"`
	DateFrom  *time.Time `json:"date_from,omitempty"`
	DateTo    *time.Time `json:"date_to,omitempty"`
	Page      *int       `json:"page,omitempty"`
	PageSize  *int       `json:"page_size,omitempty"`
	AllowedTo *string    `json:"allowed_to,omitempty"`
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

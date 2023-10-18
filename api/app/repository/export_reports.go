package repository

import (
	"context"
)

const (
	sqlExportReports = `
	with tab as (
		select id,
			   display_name,
			   date,
			   start_time,
			   end_time,
			   break_time,
			   work_time,
			   body
		from main.reports as r
		inner join main.reports_to_users rtu on r.id = rtu.report_id
		where rtu.user_id = $1 and
		    r.deleted_at is null
	)
	select  t.id                      as report_id,
       		t.display_name            as display_name,
			t.date as date,
			t.start_time as start_time,
			t.end_time as end_time,
			t.break_time as break_time,
			t.work_time as work_time,
			t.body as body
	from tab as t`
)

type ExportReports struct {
	InvokerId string `json:"invoker_id,omitempty"`
}

func (c *Client) ExportReports(ctx context.Context, msg *ExportReports) ([]ExportedReport, error) {
	row, err := c.driver.Query(ctx, sqlExportReports,
		msg.InvokerId,
	)
	if err != nil {
		return nil, NewInternalError(err)
	}

	reports := make([]ExportedReport, 0)

	for row.Next() {
		report := ExportedReport{}
		err := row.Scan(
			&report.Id,
			&report.DisplayName,
			&report.Date,
			&report.StartTime,
			&report.EndTime,
			&report.BreakTime,
			&report.WorkTime,
			&report.Body,
		)
		if err != nil {
			return nil, NewInternalError(err)
		}
		reports = append(reports, report)
	}

	return reports, nil
}

type ExportedReport struct {
	Id          *string `json:"id,omitempty"`
	DisplayName *string `json:"display_name,omitempty"`
	//Date        *time.Time `json:"date,omitempty"`
	Date      *int64  `json:"date,omitempty"`
	StartTime *int64  `json:"start_time,omitempty"`
	EndTime   *int64  `json:"end_time,omitempty"`
	BreakTime *int64  `json:"break_time,omitempty"`
	WorkTime  *int64  `json:"work_time,omitempty"`
	Body      *string `json:"body,omitempty"`
}

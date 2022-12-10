package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
	"time"
)

const sqlUpdateReport = `
	update main.reports
    set display_name = coalesce($2, display_name), 
        date = coalesce($3, date), 
        start_time = coalesce($4, start_time), 
        end_time = coalesce($5, end_time), 
        break_time = coalesce($6, break_time), 
        work_time = coalesce($7, work_time), 
        body = coalesce($8, body), 
        updated_at = now()
    where id = $1 and exists(select 1 
							 from main.reports
							 inner join main.reports_to_users rtu on reports.id = rtu.report_id
							 where rtu.user_id = $9)`

func (c *Client) UpdateReport(ctx context.Context, msg *UpdateReport) error {
	err := c.isFindReport(ctx, msg.InvokerId, msg.ReportId)
	if errors.Is(err, ErrReportIdNotFound) {
		return ErrReportIdNotFound
	}

	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})

	row, err := transaction.Query(ctx, sqlUpdateReport,
		msg.ReportId,
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
		return NewInternalError(err)
	}

	row.Next()
	status := row.CommandTag()
	if status != nil && !status.Update() {
		return NewInternalError(err)
	}

	if err != nil {
		return NewInternalError(err)
	}

	_ = transaction.Commit(ctx)

	return nil
}

type UpdateReport struct {
	InvokerId   string     `json:"invoker_id,omitempty"`
	ReportId    string     `json:"id,omitempty"`
	DisplayName *string    `json:"display_name,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	StartTime   *int64     `json:"start_time,omitempty"`
	EndTime     *int64     `json:"end_time,omitempty"`
	BreakTime   *int64     `json:"break_time,omitempty"`
	WorkTime    *int64     `json:"work_time,omitempty"`
	Body        *string    `json:"body,omitempty"`
}

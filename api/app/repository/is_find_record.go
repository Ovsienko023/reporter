package repository

import (
	"context"
)

const sqlFindReport = `
	select 1
    from main.reports
    inner join main.reports_to_users rtu on reports.id = rtu.report_id
	where id = $1 and rtu.user_id = $2`

func (c *Client) isFindReport(ctx context.Context, invokerId, reportId string) error {
	row, err := c.driver.Query(ctx, sqlFindReport, reportId, invokerId)
	if err != nil {
		return NewInternalError(err)
	}

	var report *int

	for row.Next() {
		err = row.Scan(&report)
		if err != nil {
			return NewInternalError(err)
		}
	}

	if report == nil {
		return ErrReportIdNotFound
	}

	return nil
}

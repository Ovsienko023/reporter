package database

import (
	"context"
)

const sqlFindReport = `
	select 1
    from main.reports
    where id = $1`

func (c *Client) isFindReport(ctx context.Context, reportId string) error {
	row, err := c.driver.Query(ctx, sqlFindReport, reportId)
	if err != nil {
		return NewInternalError(err)
	}

	var report *int

	row.Next()
	err = row.Scan(&report)
	if err != nil {
		return NewInternalError(err)
	}

	if report == nil {
		return ErrReportIdNotFound
	}

	return nil
}

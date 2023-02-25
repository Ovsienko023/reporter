package repository

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v4"
)

const sqlDeleteReport = `
	delete from main.reports
    where id = $1 and creator_id = $2`

func (c *Client) DeleteReport(ctx context.Context, msg *DeleteReport) error {
	err := c.isFindReport(ctx, msg.InvokerId, msg.ReportId)
	if errors.Is(err, ErrReportIdNotFound) {
		return ErrReportIdNotFound
	}

	transaction, err := c.driver.BeginTx(ctx, pgx.TxOptions{})
	row, err := transaction.Query(ctx, sqlDeleteReport,
		msg.ReportId,
		msg.InvokerId,
	)
	if err != nil {
		return NewInternalError(err)
	}

	row.Next()
	status := row.CommandTag()
	if status != nil && !status.Delete() {
		return NewInternalError(err)
	}

	if err != nil {
		return NewInternalError(err)
	}

	_ = transaction.Commit(ctx)

	return nil
}

type DeleteReport struct {
	InvokerId string `json:"invoker_id,omitempty"`
	ReportId  string `json:"id,omitempty"`
}
